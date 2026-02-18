package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"
)

type Parser struct {
	analysisResult AnalysisResult
}

func NewParser() *Parser {
	result := &Parser{}
	result.analysisResult = AnalysisResult{}
	return result
}

func (this *Parser) Run(pattern string) AnalysisResult {

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

	//为每个包运行分析器
	for _, pkg := range pkgs {
		this.runPackageAnalysis(pkg)
	}
	return this.analysisResult
}

func (this *Parser) runPackageAnalysis(pkg *packages.Package) {
	analyser := &analysis.Analyzer{
		Name:     "enhancedfunccall",
		Doc:      "找出并分类所有的类和方法",
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

func (this *Parser) getReceiverTargetType(methodInfo MethodInfo) string {
	receiverType := methodInfo.Receiver
	if receiverType.IsPointer {
		return receiverType.ElemType.Code
	} else {
		return receiverType.Code
	}
}

func (this *Parser) runAnalysis(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// 获取包的绝对路径
	pkgDir := ""
	if len(pass.Files) > 0 {
		// 从第一个文件获取包目录
		file := pass.Files[0]
		if pass.Fset != nil && file.Pos().IsValid() {
			pos := pass.Fset.Position(file.Pos())
			pkgDir = filepath.Dir(pos.Filename)
		}
	}

	pkgInfo := PackageInfo{
		Name:    pass.Pkg.Name(),
		Path:    pass.Pkg.Path(),
		Dir:     pkgDir,
		Imports: nil,
		Structs: []StructInfo{},
	}

	// 1. 提取导入信息
	this.extractImportsFromFile(pass.Files, &pkgInfo)

	// 2. 定义过滤器，只关注类型定义和方法声明
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),  // 通用声明（包括类型定义）
		(*ast.FuncDecl)(nil), // 函数声明
	}

	// 存储临时结构体信息
	structNameToStructInfo := map[string]StructInfo{}

	// 3. 遍历AST节点
	inspect.WithStack(nodeFilter, func(n ast.Node, push bool, stack []ast.Node) bool {
		if !push {
			return true
		}

		switch node := n.(type) {
		case *ast.GenDecl:
			if node.Tok == token.TYPE {
				for _, spec := range node.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							// 提取结构体信息
							structInfo := this.extractStructInfo(pass, node, typeSpec, structType, pkgInfo.Imports)
							oldStruct, isExist := structNameToStructInfo[structInfo.Name]
							if isExist == false {
								//不存在的时候直接赋值
								structNameToStructInfo[structInfo.Name] = structInfo
							} else {
								//存在的时候，赋值methods
								structInfo.Methods = oldStruct.Methods
								structNameToStructInfo[structInfo.Name] = structInfo
							}
						}
					}
				}
			}

		case *ast.FuncDecl:
			if node.Recv != nil {
				// 这是一个方法
				methodInfo := this.extractMethodInfo(pass, node, pkgInfo.Imports)

				structInfoName := this.getReceiverTargetType(methodInfo)
				oldStruct, isExist := structNameToStructInfo[structInfoName]
				if isExist == false {
					//不存在的时候直接赋值
					structNameToStructInfo[structInfoName] = StructInfo{
						Methods: []MethodInfo{methodInfo},
					}
				} else {
					//存在的时候，直接添加methods
					oldStruct.Methods = append(oldStruct.Methods, methodInfo)
					structNameToStructInfo[structInfoName] = oldStruct
				}
			}
		}

		return true
	})

	pkgInfo.Structs = []StructInfo{}
	for _, structInfo := range structNameToStructInfo {
		pkgInfo.Structs = append(pkgInfo.Structs, structInfo)
	}
	this.analysisResult.Packages = append(this.analysisResult.Packages, pkgInfo)

	return nil, nil
}

// extractImportsFromFile 从文件中提取导入信息
func (this *Parser) extractImportsFromFile(files []*ast.File, pkgInfo *PackageInfo) {
	result := []ImportInfo{}
	for _, file := range files {
		for _, imp := range file.Imports {
			importPath := strings.Trim(imp.Path.Value, "\"")
			var info ImportInfo
			info.PkgPath = importPath

			// 提取包名
			pkgName := filepath.Base(importPath)

			if imp.Name == nil {
				// 原名导入
				info.PkgName = pkgName
				info.ImportType = importTypeOriginal
				info.Code = fmt.Sprintf("\"%s\"", importPath)
			} else if imp.Name.Name == "." {
				// 点号导入
				info.PkgName = "."
				info.ImportType = importTypeDot
				info.Code = fmt.Sprintf(". \"%s\"", importPath)
			} else if imp.Name.Name == "_" {
				// 匿名导入
				continue
			} else if imp.Name.Name != pkgName {
				// 重命名导入
				info.PkgName = imp.Name.Name
				info.ImportType = importTypeAlias
				info.Code = fmt.Sprintf("%s \"%s\"", imp.Name.Name, importPath)
			} else {
				// 原名导入
				info.PkgName = pkgName
				info.ImportType = importTypeOriginal
				info.Code = fmt.Sprintf("\"%s\"", importPath)
			}
			result = append(result, info)
		}
	}
	pkgInfo.Imports = NewPackageImportFinder(result)
}

// extractStructInfo 提取结构体信息
func (this *Parser) extractStructInfo(pass *analysis.Pass, genDecl *ast.GenDecl,
	typeSpec *ast.TypeSpec, structType *ast.StructType, imports *PackageImportFinder) StructInfo {

	file := pass.Fset.File(genDecl.Pos()).Name()
	line := pass.Fset.Position(genDecl.Pos()).Line

	structInfo := StructInfo{
		Name:     typeSpec.Name.Name,
		Package:  pass.Pkg.Path(),
		File:     file,
		Line:     line,
		Methods:  []MethodInfo{},
		Fields:   []FieldInfo{},
		IsPublic: isPublicIdentifier(typeSpec.Name.Name),
	}

	// 提取字段信息
	if structType.Fields != nil {
		for _, field := range structType.Fields.List {
			depImports := map[string]ImportInfo{}
			fieldType := this.extractTypeFromExpr(pass, field.Type, pass.Pkg.Path(), imports, depImports)
			for _, name := range field.Names {
				fieldInfo := FieldInfo{
					Name:     name.Name,
					Type:     fieldType,
					IsPublic: isPublicIdentifier(name.Name),
				}

				// 提取标签
				if field.Tag != nil {
					fieldInfo.Tag = field.Tag.Value
				}

				structInfo.Fields = append(structInfo.Fields, fieldInfo)
			}
		}
	}

	return structInfo
}

func (this *Parser) combineDepImports(allDepImports map[string]ImportInfo, depImports map[string]ImportInfo) {
	for pkgPath, importInfo := range depImports {
		allDepImports[pkgPath] = importInfo
	}
}

// extractMethodInfo 提取方法信息
func (this *Parser) extractMethodInfo(pass *analysis.Pass, funcDecl *ast.FuncDecl, imports *PackageImportFinder) MethodInfo {
	file := pass.Fset.File(funcDecl.Pos()).Name()
	line := pass.Fset.Position(funcDecl.Pos()).Line

	method := MethodInfo{
		Name:       funcDecl.Name.Name,
		File:       file,
		Line:       line,
		Params:     []FunctionParam{},
		Results:    []FunctionParam{},
		DepImports: []ImportInfo{},
		IsPublic:   isPublicIdentifier(funcDecl.Name.Name),
	}

	allDepImports := map[string]ImportInfo{}

	// 提取接收器
	if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
		depImports := map[string]ImportInfo{}
		receiverType := this.extractTypeFromExpr(pass, funcDecl.Recv.List[0].Type, pass.Pkg.Path(), imports, depImports)
		method.Receiver = receiverType

		this.combineDepImports(allDepImports, depImports)
	}

	// 提取参数
	if funcDecl.Type.Params != nil {
		paramIndex := 0
		for _, field := range funcDecl.Type.Params.List {
			depImports := map[string]ImportInfo{}
			paramType := this.extractTypeFromExpr(pass, field.Type, pass.Pkg.Path(), imports, depImports)

			if len(field.Names) > 0 {
				for _, name := range field.Names {
					param := FunctionParam{
						Name: name.Name,
						Type: paramType,
						Code: fmt.Sprintf("%s %s", name.Name, paramType.Code),
					}
					method.Params = append(method.Params, param)
				}
			} else {
				// 匿名参数
				param := FunctionParam{
					Name: "",
					Type: paramType,
					Code: fmt.Sprintf("%s", paramType.Code),
				}
				method.Params = append(method.Params, param)
				paramIndex++
			}

			this.combineDepImports(allDepImports, depImports)
		}
	}

	// 提取返回值
	if funcDecl.Type.Results != nil {
		resultIndex := 0
		for _, field := range funcDecl.Type.Results.List {
			depImports := map[string]ImportInfo{}
			resultType := this.extractTypeFromExpr(pass, field.Type, pass.Pkg.Path(), imports, depImports)

			if len(field.Names) > 0 {
				for _, name := range field.Names {
					result := FunctionParam{
						Name: name.Name,
						Type: resultType,
						Code: fmt.Sprintf("%s %s", name.Name, resultType.Code),
					}
					method.Results = append(method.Results, result)
				}
			} else {
				// 匿名返回值
				result := FunctionParam{
					Name: "",
					Type: resultType,
					Code: fmt.Sprintf("%s", resultType.Code),
				}
				method.Results = append(method.Results, result)
				resultIndex++
			}

			this.combineDepImports(allDepImports, depImports)
		}
	}

	method.DepImports = this.flatDepImports(allDepImports)

	return method
}

func (this *Parser) flatDepImports(depImports map[string]ImportInfo) []ImportInfo {
	result := make([]ImportInfo, len(depImports), len(depImports))
	index := 0
	for _, importInfo := range depImports {
		result[index] = importInfo
		index++
	}
	return result
}

func (this *Parser) extractFieldListTypeCode(pass *analysis.Pass, fieldList *ast.FieldList, currentPkg string, imports *PackageImportFinder, depImports map[string]ImportInfo) string {
	var result = []string{}
	if fieldList == nil {
		return ""
	}
	for _, singleField := range fieldList.List {
		typeInfo := this.extractTypeFromExpr(pass, singleField.Type, currentPkg, imports, depImports)
		if singleField.Names != nil {
			for _, singleName := range singleField.Names {
				result = append(result, fmt.Sprintf("%s %s", singleName.Name, typeInfo.Code))
			}
		} else {
			result = append(result, fmt.Sprintf("%s", typeInfo.Code))
		}
	}
	return strings.Join(result, ",")
}

// extractTypeFromExpr 从表达式提取类型信息
func (this *Parser) extractTypeFromExpr(pass *analysis.Pass, expr ast.Expr, currentPkg string, imports *PackageImportFinder, depImports map[string]ImportInfo) TypeInfo {

	var typeInfo TypeInfo

	switch t := expr.(type) {
	case *ast.Ident:
		// 标识符类型
		typeInfo.Code = t.Name

		// 检查是否是内置类型
		if isBuiltinType(t.Name) {
			typeInfo.Package = "_builtin_"
			typeInfo.DepImports = nil
		} else {
			// 尝试从类型信息中获取
			if pass.TypesInfo != nil {
				if obj := pass.TypesInfo.ObjectOf(t); obj != nil {
					if pkg := obj.Pkg(); pkg != nil {
						typeInfo.Package = pkg.Path()
						importInfo, isExists := imports.GetImportByPath(pkg.Path())
						if isExists {
							depImports[pkg.Path()] = importInfo
							typeInfo.DepImports = []ImportInfo{importInfo}
						}
					}
				}
			}

			// 如果没找到，假设是当前包的类型
			if typeInfo.Package == "" {
				typeInfo.Package = currentPkg
				typeInfo.DepImports = nil
			}
		}

	case *ast.SelectorExpr:
		// 选择器表达式: package.Type
		if x, ok := t.X.(*ast.Ident); ok {
			pkgName := x.Name
			typeName := t.Sel.Name

			typeInfo.Code = pkgName + "." + typeName

			// 从导入信息中查找
			innerType := this.extractTypeFromExpr(pass, t.Sel, currentPkg, imports, depImports)
			typeInfo.Package = innerType.Package
			typeInfo.DepImports = this.flatDepImports(depImports)
		}

	case *ast.StarExpr:
		// 指针类型
		innerType := this.extractTypeFromExpr(pass, t.X, currentPkg, imports, depImports)
		typeInfo = TypeInfo{
			Code:       "*" + innerType.Code,
			Package:    innerType.Package,
			IsPointer:  true,
			ElemType:   &innerType,
			DepImports: this.flatDepImports(depImports),
		}

	case *ast.ArrayType:
		// 数组或切片类型
		if t.Len == nil {
			// 切片类型
			innerType := this.extractTypeFromExpr(pass, t.Elt, currentPkg, imports, depImports)
			typeInfo = TypeInfo{
				Code:       "[]" + innerType.Code,
				Package:    innerType.Package,
				IsSlice:    true,
				ElemType:   &innerType,
				DepImports: this.flatDepImports(depImports),
			}
		} else {
			// 数组类型
			innerType := this.extractTypeFromExpr(pass, t.Elt, currentPkg, imports, depImports)
			typeInfo = TypeInfo{
				Code:       fmt.Sprintf("[%v]%s", t.Len, innerType.Code),
				Package:    innerType.Package,
				IsArray:    true,
				DepImports: this.flatDepImports(depImports),
			}
		}

	case *ast.MapType:
		// Map类型
		keyType := this.extractTypeFromExpr(pass, t.Key, currentPkg, imports, depImports)
		valueType := this.extractTypeFromExpr(pass, t.Value, currentPkg, imports, depImports)
		typeInfo = TypeInfo{
			Code:       fmt.Sprintf("map[%s]%s", keyType.Code, valueType.Code),
			Package:    currentPkg,
			IsMap:      true,
			MapKey:     &keyType,
			MapValue:   &valueType,
			DepImports: this.flatDepImports(depImports),
		}
	case *ast.Ellipsis:
		//ellipse
		elemType := this.extractTypeFromExpr(pass, t.Elt, currentPkg, imports, depImports)
		typeInfo = TypeInfo{
			Code:       fmt.Sprintf("...%s", elemType.Code),
			Package:    currentPkg,
			IsEllipse:  true,
			ElemType:   &elemType,
			DepImports: this.flatDepImports(depImports),
		}
	case *ast.FuncType:
		//函数类型
		paramCode := this.extractFieldListTypeCode(pass, t.Params, currentPkg, imports, depImports)
		resultCode := this.extractFieldListTypeCode(pass, t.Results, currentPkg, imports, depImports)
		typeInfo = TypeInfo{
			Code:       fmt.Sprintf("func(%s)(%s)", paramCode, resultCode),
			Package:    currentPkg,
			IsFunc:     true,
			ElemType:   nil,
			DepImports: this.flatDepImports(depImports),
		}
	case *ast.InterfaceType:
		//接口类型
		code := strings.Builder{}
		for _, singleMethod := range t.Methods.List {
			fieldListInner := []*ast.Field{singleMethod}
			fieldList := &ast.FieldList{List: fieldListInner}
			code.WriteString("\n")
			code.WriteString(this.extractFieldListTypeCode(pass, fieldList, currentPkg, imports, depImports))
		}
		typeInfo = TypeInfo{
			Code:        fmt.Sprintf("interface{%s}\n", code),
			Package:     currentPkg,
			IsInterface: true,
			ElemType:    nil,
			DepImports:  this.flatDepImports(depImports),
		}
	case *ast.ChanType:
		//chan类型
		elemType := this.extractTypeFromExpr(pass, t.Value, currentPkg, imports, depImports)
		typeInfo = TypeInfo{
			Code:       fmt.Sprintf("chan ", elemType.Code),
			Package:    currentPkg,
			IsChan:     true,
			ElemType:   &elemType,
			DepImports: this.flatDepImports(depImports),
		}
	default:
		panic(fmt.Sprintf("%T unknown fieldType ", t))
	}

	return typeInfo
}

// isPublicIdentifier 检查标识符是否为公开（首字母大写）
func isPublicIdentifier(name string) bool {
	if len(name) == 0 {
		return false
	}
	firstRune := []rune(name)[0]
	return unicode.IsUpper(firstRune)
}

// isBuiltinType 检查是否是内置类型
func isBuiltinType(name string) bool {
	builtinTypes := map[string]bool{
		"bool":       true,
		"byte":       true,
		"complex64":  true,
		"complex128": true,
		"error":      true,
		"float32":    true,
		"float64":    true,
		"int":        true,
		"int8":       true,
		"int16":      true,
		"int32":      true,
		"int64":      true,
		"rune":       true,
		"string":     true,
		"uint":       true,
		"uint8":      true,
		"uint16":     true,
		"uint32":     true,
		"uint64":     true,
		"uintptr":    true,
		"any":        true,
	}
	return builtinTypes[name]
}

type ImportTypeEnum struct {
	name string
}

var (
	importTypeOriginal       = ImportTypeEnum{name: "original"}
	importTypeDot            = ImportTypeEnum{name: "dot"}
	importTypeAlias          = ImportTypeEnum{name: "alias"}
	importTypeCurrentPackage = ImportTypeEnum{name: "current_package"}
)

// ImportInfo 记录导入包的信息
type ImportInfo struct {
	Code       string         // 在代码中使用导入的代码
	PkgPath    string         // 完整的导入路径
	PkgName    string         // 在代码中使用的名称
	ImportType ImportTypeEnum // 导入方式: "original", "dot", "alias", "current_package"
}

// TypeInfo 记录类型信息
type TypeInfo struct {
	Code        string       // 类型名称
	Package     string       // 包路径
	IsPointer   bool         // 是否为指针类型
	IsSlice     bool         // 是否为切片类型
	IsArray     bool         // 是否为数组类型
	IsMap       bool         // 是否为Map类型
	IsInterface bool         //是否为Interface类型
	IsFunc      bool         //是否为Func类型
	IsChan      bool         //是否为Chan类型
	IsEllipse   bool         //是否为省略类型
	MapKey      *TypeInfo    // Map的键类型
	MapValue    *TypeInfo    // Map的值类型
	ElemType    *TypeInfo    // 指针/切片/数组的元素类型
	DepImports  []ImportInfo // 导入信息
}

// FunctionParam 记录函数参数信息
type FunctionParam struct {
	Code string
	Name string
	Type TypeInfo
}

// MethodInfo 记录方法信息
type MethodInfo struct {
	Name       string
	Receiver   TypeInfo
	Params     []FunctionParam
	Results    []FunctionParam
	DepImports []ImportInfo // 依赖的所有类型
	File       string       // 方法所在的文件
	Line       int          // 方法所在行
	IsPublic   bool
}

// StructInfo 记录结构体信息
type StructInfo struct {
	Name     string
	Package  string
	File     string
	Line     int
	Methods  []MethodInfo
	Fields   []FieldInfo // 结构体字段
	IsPublic bool
}

// FieldInfo 记录结构体字段信息
type FieldInfo struct {
	Name     string
	Type     TypeInfo
	Tag      string // 标签
	IsPublic bool
}

// PackageInfo 记录包信息
type PackageInfo struct {
	Name    string
	Path    string
	Dir     string
	Imports *PackageImportFinder
	Structs []StructInfo
}

// AnalysisResult 分析结果
type AnalysisResult struct {
	Packages []PackageInfo
}

type PackageImportFinder struct {
	imports             []ImportInfo
	packagePathToImport map[string]ImportInfo
	packageNameToImport map[string]ImportInfo
}

func NewPackageImportFinder(imports []ImportInfo) *PackageImportFinder {
	result := &PackageImportFinder{}
	//创建pkgPath到import的映射
	result.packagePathToImport = map[string]ImportInfo{}
	result.packageNameToImport = map[string]ImportInfo{}
	for _, ipt := range imports {
		result.packagePathToImport[ipt.PkgPath] = ipt
		result.packageNameToImport[ipt.PkgName] = ipt
	}

	//利用packagePathToImport去重，重新得到imports
	result.imports = []ImportInfo{}
	for _, ipt := range result.packagePathToImport {
		result.imports = append(result.imports, ipt)
	}
	return result
}

func (this *PackageImportFinder) GetImportByPath(path string) (ImportInfo, bool) {
	result, isExist := this.packagePathToImport[path]
	return result, isExist
}

func (this *PackageImportFinder) GetImportByName(name string) (ImportInfo, bool) {
	result, isExist := this.packagePathToImport[name]
	return result, isExist
}
