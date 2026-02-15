package main

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"
)

type Parser struct {
	workingDir      string
	initPackageName string
	initPackagePath string
	walker          *func(*CallInfo)
}

func NewParser() *Parser {
	result := &Parser{
		workingDir:      "",
		initPackageName: "",
		initPackagePath: "",
		walker:          nil,
	}

	// 在 init 中获取运行目录
	var err error
	result.workingDir, err = os.Getwd()
	if err != nil {
		panic("获取当前目录失败")
	}

	return result
}

func (this *Parser) Run(pattern string, walker func(*CallInfo)) {

	this.walker = &walker

	// 配置包加载
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedSyntax,
	}

	//加载包
	pkgs, err := packages.Load(cfg, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "加载包失败: %v\n", err)
		os.Exit(1)
	}

	//设置initPackage
	for _, pkg := range pkgs {
		this.setInitPackageName(pkg)
	}

	if this.initPackageName == "" {
		panic("找不到当前工作目录的package name")
	}
	if this.initPackagePath == "" {
		panic("找不到当前工作目录的package path")
	}
	globalGeneratePackagePath = this.initPackagePath

	//为每个包运行分析器
	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			fmt.Fprintf(os.Stderr, "包 %s 有错误: %v\n", pkg.ID, pkg.Errors)
			continue
		}
		this.runPackageAnalysis(pkg)
	}

	this.walker = nil
}

func (this *Parser) runPackageAnalysis(pkg *packages.Package) {
	analyser := &analysis.Analyzer{
		Name:     "enhancedfunccall",
		Doc:      "找出并分类所有的函数调用",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      this.runAnalysis,
	}
	pass := &analysis.Pass{
		Analyzer:  analyser,
		Fset:      pkg.Fset,
		Files:     pkg.Syntax,
		TypesInfo: pkg.TypesInfo,
		Pkg:       pkg.Types,
		ResultOf:  make(map[*analysis.Analyzer]interface{}),
		Report:    func(d analysis.Diagnostic) {},
	}

	// 运行依赖的分析器 (inspect)
	inspectResult, err := inspect.Analyzer.Run(pass)
	if err != nil {
		fmt.Fprintf(os.Stderr, "运行 inspect 分析器失败: %v\n", err)
		return
	}
	pass.ResultOf[inspect.Analyzer] = inspectResult

	// 运行自定义分析器
	nil, err := analyser.Run(pass)
	if err != nil {
		fmt.Fprintf(os.Stderr, "运行参数分析器失败: %v\n", err)
		return
	}
}

func (this *Parser) setInitPackageName(pkg *packages.Package) {
	if this.initPackageName != "" {
		return
	}
	// 获取当前分析包的信息
	pkgName := pkg.Name
	pkgPath := pkg.PkgPath
	pkgDir := pkg.Dir

	// 计算相对于运行目录的路径
	if rel, err := filepath.Rel(this.workingDir, pkgDir); err == nil {
		if rel == "." {
			//找到当前目录的所在包名
			this.initPackageName = pkgName
			this.initPackagePath = pkgPath
		}
	}
}

func (this *Parser) runAnalysis(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		callExpr := n.(*ast.CallExpr)
		callInfo := analyzeFunctionCall(pass, callExpr)
		if this.walker != nil {
			walkerTarget := *this.walker
			walkerTarget(callInfo)
		}
	})
	return nil, nil
}

// ParamInfo 参数信息结构体
type ParamInfo struct {
	Index        int                // 参数索引 (0-based)
	Position     token.Position     // 位置信息
	ExprType     string             // 表达式类型 (Ident, BasicLit, CallExpr等)
	Type         types.Type         // Go类型信息
	ConstValue   interface{}        // 常量值 (如果有)
	IsConst      bool               // 是否为常量
	IsNil        bool               // 是否为nil
	IsVariadic   bool               // 是否为可变参数
	ExpectedType types.Type         // 期望的类型 (从函数签名)
	ExpectedName string             // 期望的参数名 (从函数签名)
	SourceCode   string             // 源代码片段
	Detailed     *DetailedParamInfo // 详细信息
}

func NewConstantStringParamInfo(value string) ParamInfo {
	return ParamInfo{
		Index:        0,
		Position:     token.NewFileSet().Position(token.NoPos),
		ExprType:     "Ident",
		Type:         &types.Basic{},
		ConstValue:   value,
		IsConst:      true,
		IsNil:        false,
		IsVariadic:   false,
		ExpectedType: &types.Basic{},
		ExpectedName: "",
		SourceCode:   "",
		Detailed:     nil,
	}
}

// DetailedParamInfo 详细参数信息
type DetailedParamInfo struct {
	IdentName     string           // 标识符名称 (如果是Ident)
	PkgPath       string           // 包路径 (如果有)
	IsBuiltin     bool             // 是否为内置函数/类型
	IsMethod      bool             // 是否为方法调用
	Receiver      string           // 接收器类型 (如果是方法)
	ElementType   types.Type       // 元素类型 (数组/切片/映射)
	KeyType       types.Type       // 键类型 (映射)
	InterfaceType *types.Interface // 接口类型 (如果有)
	NamedType     *types.Named     // 命名类型 (如果有)
	Signature     *types.Signature // 函数签名 (如果是函数类型)
}

// CallInfo 函数调用信息结构体
type CallInfo struct {
	FuncName    string           // 函数名
	Position    token.Position   // 调用位置
	PackageName string           // 包名
	PackagePath string           // 包路径
	IsMethod    bool             // 是否为方法
	Receiver    string           // 接收器类型
	Params      []ParamInfo      // 参数列表
	ReturnType  types.Type       // 返回值类型
	IsVariadic  bool             // 是否为可变参数函数
	Signature   *types.Signature // 函数签名
}

func analyzeFunctionCall(pass *analysis.Pass, callExpr *ast.CallExpr) *CallInfo {
	pos := pass.Fset.Position(callExpr.Pos())
	funcName := getFunctionName(callExpr.Fun)

	callInfo := &CallInfo{
		FuncName: funcName,
		Position: pos,
		Params:   make([]ParamInfo, 0, len(callExpr.Args)),
	}

	// 获取函数对象和签名
	var funcObj types.Object
	var signature *types.Signature

	switch fun := callExpr.Fun.(type) {
	case *ast.Ident:
		funcObj = pass.TypesInfo.Uses[fun]
		if funcObj != nil {
			callInfo.PackageName = getPackageNameFromObject(funcObj)
			callInfo.PackagePath = getPackagePathFromObject(funcObj)
		}
	case *ast.SelectorExpr:
		funcObj = pass.TypesInfo.Uses[fun.Sel]
		if funcObj != nil {
			callInfo.PackageName = getPackageNameFromObject(funcObj)
			callInfo.PackagePath = getPackagePathFromObject(funcObj)

			// 检查是否为方法
			if sig, ok := funcObj.Type().(*types.Signature); ok && sig.Recv() != nil {
				callInfo.IsMethod = true
				callInfo.Receiver = sig.Recv().Type().String()
			}
		}
	}

	// 获取函数签名
	if funcObj != nil {
		if sig, ok := funcObj.Type().(*types.Signature); ok {
			signature = sig
			callInfo.Signature = sig
			callInfo.IsVariadic = sig.Variadic()

			// 获取返回类型
			if sig.Results() != nil && sig.Results().Len() > 0 {
				callInfo.ReturnType = sig.Results().At(0).Type()
			}
		}
	}

	// 提取每个参数的信息
	for i, arg := range callExpr.Args {
		paramInfo := extractParameterInfo(pass, arg, signature, i)
		paramInfo.Position = pass.Fset.Position(arg.Pos())
		paramInfo.Index = i
		paramInfo.SourceCode = getSourceCodeSnippet(pass, arg)

		callInfo.Params = append(callInfo.Params, paramInfo)
	}

	return callInfo
}

func extractParameterInfo(pass *analysis.Pass, arg ast.Expr,
	signature *types.Signature, index int) ParamInfo {

	info := ParamInfo{
		ExprType: getExprTypeName(arg),
	}

	// 1. 获取类型信息
	if tv, ok := pass.TypesInfo.Types[arg]; ok {
		info.Type = tv.Type
		info.IsConst = tv.Value != nil

		if tv.Value != nil {
			// 尝试获取常量值
			info.ConstValue = getConstantValue(tv.Value)
		}

		// 检查是否为nil
		if tv.IsNil() {
			info.IsNil = true
		}

		// 设置详细信息
		info.Detailed = extractDetailedInfo(pass, arg, tv.Type)
	}

	// 2. 检查是否为可变参数
	if signature != nil && signature.Variadic() {
		params := signature.Params()
		if index >= params.Len()-1 {
			// 最后一个参数是可变参数
			info.IsVariadic = true
		}
	}

	// 3. 获取期望的类型和参数名（如果有函数签名）
	if signature != nil && signature.Params() != nil && index < signature.Params().Len() {
		param := signature.Params().At(index)
		info.ExpectedType = param.Type()
		info.ExpectedName = param.Name()
	}

	return info
}

func extractDetailedInfo(pass *analysis.Pass, arg ast.Expr, typ types.Type) *DetailedParamInfo {
	detailed := &DetailedParamInfo{}

	// 处理标识符
	if ident, ok := arg.(*ast.Ident); ok {
		detailed.IdentName = ident.Name
		if obj := pass.TypesInfo.Uses[ident]; obj != nil {
			if pkg := obj.Pkg(); pkg != nil {
				detailed.PkgPath = pkg.Path()
			}
			if obj.Parent() == types.Universe {
				detailed.IsBuiltin = true
			}
		}
	}

	// 处理选择器表达式（如 pkg.Func 或 var.Method）
	if sel, ok := arg.(*ast.SelectorExpr); ok {
		if obj := pass.TypesInfo.Uses[sel.Sel]; obj != nil {
			// 检查是否是方法
			if sig, ok := obj.Type().(*types.Signature); ok && sig.Recv() != nil {
				detailed.IsMethod = true
				detailed.Receiver = sig.Recv().Type().String()
			}
		}
	}

	// 类型详细信息
	switch t := typ.(type) {
	case *types.Named:
		detailed.NamedType = t
		if pkg := t.Obj().Pkg(); pkg != nil {
			detailed.PkgPath = pkg.Path()
		}

	case *types.Interface:
		detailed.InterfaceType = t

	case *types.Signature:
		detailed.Signature = t

	case *types.Slice:
		detailed.ElementType = t.Elem()

	case *types.Array:
		detailed.ElementType = t.Elem()

	case *types.Map:
		detailed.KeyType = t.Key()
		detailed.ElementType = t.Elem()

	case *types.Chan:
		detailed.ElementType = t.Elem()
	}

	return detailed
}

func getExprTypeName(expr ast.Expr) string {
	switch expr.(type) {
	case *ast.Ident:
		return "Ident"
	case *ast.BasicLit:
		return "BasicLit"
	case *ast.CompositeLit:
		return "CompositeLit"
	case *ast.CallExpr:
		return "CallExpr"
	case *ast.SelectorExpr:
		return "SelectorExpr"
	case *ast.IndexExpr:
		return "IndexExpr"
	case *ast.SliceExpr:
		return "SliceExpr"
	case *ast.TypeAssertExpr:
		return "TypeAssertExpr"
	case *ast.UnaryExpr:
		return "UnaryExpr"
	case *ast.BinaryExpr:
		return "BinaryExpr"
	case *ast.FuncLit:
		return "FuncLit"
	case *ast.StarExpr:
		return "StarExpr"
	default:
		return fmt.Sprintf("%T", expr)
	}
}

func getFunctionName(fun ast.Expr) string {
	switch f := fun.(type) {
	case *ast.Ident:
		return f.Name
	case *ast.SelectorExpr:
		if _, ok := f.X.(*ast.Ident); ok {
			return f.Sel.Name
		}
	}
	return "unknown"
}

func getConstantValue(value constant.Value) interface{} {
	if value == nil {
		return nil
	}

	switch value.Kind() {
	case constant.Bool:
		return constant.BoolVal(value)
	case constant.String:
		return constant.StringVal(value)
	case constant.Int:
		if v, ok := constant.Int64Val(value); ok {
			return v
		}
		return value.ExactString()
	case constant.Float:
		if v, ok := constant.Float64Val(value); ok {
			return v
		}
		return value.ExactString()
	case constant.Complex:
		return value.ExactString()
	default:
		return value.ExactString()
	}
}

func getSourceCodeSnippet(pass *analysis.Pass, expr ast.Expr) string {
	start := pass.Fset.Position(expr.Pos())
	end := pass.Fset.Position(expr.End())

	// 简单实现：返回位置信息
	// 实际实现可以读取文件内容
	return fmt.Sprintf("%s:%d-%d", start.Filename, start.Line, end.Line)
}

func getPackageNameFromObject(obj types.Object) string {
	if pkg := obj.Pkg(); pkg != nil {
		return pkg.Name()
	}
	return ""
}

func getPackagePathFromObject(obj types.Object) string {
	if pkg := obj.Pkg(); pkg != nil {
		return pkg.Path()
	}
	return ""
}
