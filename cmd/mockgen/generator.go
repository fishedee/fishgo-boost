package main

import (
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Generator struct {
	typeReg *regexp.Regexp
	fileReg *regexp.Regexp
}

func NewGenerator(fileRegex string, typeRegex string) *Generator {
	var err error
	result := &Generator{}
	result.fileReg, err = regexp.CompilePOSIX(fileRegex)
	if err != nil {
		panic(fmt.Sprintf("编译文件正则表达式[%s]失败", fileRegex))
	}
	result.typeReg, err = regexp.CompilePOSIX(typeRegex)
	if err != nil {
		panic(fmt.Sprintf("编译类型正则表达式[%s]失败", typeRegex))
	}
	return result
}

func (this *Generator) Run(analysisResult AnalysisResult) error {

	for _, pkg := range analysisResult.Packages {
		pkg = this.filterContent(pkg)

		if len(pkg.Structs) == 0 {
			continue
		}
		result, err := this.generateContent(pkg)
		if err != nil {
			panic(err)
		}

		result, err = this.formatSource(result)
		if err != nil {
			panic(err)
		}

		err = this.writeFile(pkg.Dir, result)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (this *Generator) generateContent(pkg PackageInfo) (string, error) {
	packageInfo := "package " + pkg.Name + "\n"

	contentInfo, err := this.generateSingleFileCode(pkg)
	if err != nil {
		return "", err
	}

	importInfo, err := this.generateSingleFileImport(pkg)
	if err != nil {
		return "", err
	}
	return packageInfo + importInfo + contentInfo, nil
}

func (this *Generator) combineDepImports(allDepImports map[string]ImportInfo, depImports []ImportInfo) {
	for _, importInfo := range depImports {
		allDepImports[importInfo.PkgPath] = importInfo
	}
}

func (this *Generator) generateSingleFileImport(pkg PackageInfo) (string, error) {
	depImports := map[string]ImportInfo{}
	for _, structInfo := range pkg.Structs {
		for _, methodInfo := range structInfo.Methods {
			this.combineDepImports(depImports, methodInfo.DepImports)
		}
	}
	proxyPath := "github.com/fishedee/fishgo-boost/app/proxy"
	_, isExists := depImports[proxyPath]
	if isExists == false {
		depImports[proxyPath] = ImportInfo{
			PkgName:    ".",
			PkgPath:    proxyPath,
			Code:       ". \"" + proxyPath + "\"",
			ImportType: importTypeDot,
		}
	}

	allImports := []string{}
	for _, depImport := range depImports {
		allImports = append(allImports, depImport.Code)
	}
	return fmt.Sprintf("import ( %s )\n", strings.Join(allImports, "\n")), nil
}

func (this *Generator) generateSingleFileCode(pkg PackageInfo) (string, error) {
	//类型排序
	sort.Slice(pkg.Structs, func(i int, j int) bool {
		return pkg.Structs[i].Name < pkg.Structs[j].Name
	})
	typeCodes := []string{}
	for _, structInfo := range pkg.Structs {
		//方法排序
		sort.Slice(structInfo.Methods, func(i int, j int) bool {
			return structInfo.Methods[i].Name < structInfo.Methods[j].Name
		})
		typeCodes = append(typeCodes, this.generateSingleTypeInterface(structInfo))
		typeCodes = append(typeCodes, this.generateSingleTypeMock(structInfo))
	}

	mockImplements := strings.Join(typeCodes, "\n")

	initCodes := []string{}
	for _, singleTypeInfo := range pkg.Structs {
		initCode := this.generateSingleInit(singleTypeInfo)
		initCodes = append(initCodes, initCode)
	}

	initImplements := fmt.Sprintf("func init(){\n %s \n}\n", strings.Join(initCodes, "\n"))

	return mockImplements + "\n" + initImplements, nil
}

func (this *Generator) filterContent(pkg PackageInfo) PackageInfo {
	newStructs := []StructInfo{}
	for _, structInfo := range pkg.Structs {
		if structInfo.IsPublic == false {
			continue
		}
		if this.fileReg.Match([]byte(structInfo.File)) == false {
			continue
		}
		if this.typeReg.Match([]byte(structInfo.Name)) == false {
			continue
		}
		newMethods := []MethodInfo{}
		for _, methodInfo := range structInfo.Methods {
			if methodInfo.IsPublic == false {
				continue
			}
			newMethods = append(newMethods, methodInfo)
		}
		structInfo.Methods = newMethods
		newStructs = append(newStructs, structInfo)
	}

	result := pkg
	result.Structs = newStructs
	return result
}

func (this *Generator) formatSource(data string) (string, error) {
	result, err := format.Source([]byte(data))
	if err != nil {
		return "", errors.New(err.Error() + "," + data)
	}
	return string(result), nil
}

func (this *Generator) writeFile(dir string, data string) error {
	basePath := filepath.Base(dir)
	filePath := filepath.Join(dir, basePath+"_mock.go")
	oldData, err := os.ReadFile(filePath)
	if err == nil && string(oldData) == data {
		return nil
	}
	return os.WriteFile(filePath, []byte(data), 0644)
}

func (this *Generator) getInterfaceName(typeName string) string {
	return "I" + strings.ToUpper(typeName[0:1]) + typeName[1:]
}

func (this *Generator) getMockName(typeName string) string {
	return typeName + "Mock"
}

func (this *Generator) generateSingleField(data []FunctionParam) string {
	var result []string
	for _, singleData := range data {
		result = append(result, singleData.Code)
	}
	return strings.Join(result, ",")
}

func (this *Generator) generateSingleFieldName(data []FunctionParam) string {
	var result []string
	for _, singleData := range data {
		result = append(result, singleData.Name)
	}
	return strings.Join(result, ",")
}

func (this *Generator) generateSingleResult(data []FunctionParam) string {
	var result []string
	for _, singleData := range data {
		result = append(result, singleData.Type.Code)
	}
	return strings.Join(result, ",")
}

func (this *Generator) generateSingleTypeInterface(structInfo StructInfo) string {
	funResult := []string{}
	for _, method := range structInfo.Methods {
		single := fmt.Sprintf("%s(%s)(%s)",
			method.Name,
			this.generateSingleField(method.Params),
			this.generateSingleResult(method.Results),
		)
		funResult = append(funResult, single)
	}
	return fmt.Sprintf("type %s interface{\n %s \n}\n",
		this.getInterfaceName(structInfo.Name),
		strings.Join(funResult, "\n"),
	)
}

func (this *Generator) generateSingleTypeMock(structInfo StructInfo) string {
	funResult := []string{}
	for _, method := range structInfo.Methods {
		single := fmt.Sprintf("%s func (%s)(%s)",
			method.Name+"Handler",
			this.generateSingleField(method.Params),
			this.generateSingleResult(method.Results))
		funResult = append(funResult, single)
	}
	mockTypeStruct := fmt.Sprintf("type %s struct{\n %s \n}\n",
		this.getMockName(structInfo.Name),
		strings.Join(funResult, "\n"),
	)

	methodResult := []string{}
	for _, method := range structInfo.Methods {
		returnCode := ""
		if len(method.Results) > 0 {
			returnCode = "return"
		}
		single := fmt.Sprintf("func (this * %s) %s (%s)(%s){\n %s this.%s(%s)\n}\n",
			this.getMockName(structInfo.Name),
			method.Name,
			this.generateSingleField(method.Params),
			this.generateSingleResult(method.Results),
			returnCode,
			method.Name+"Handler",
			this.generateSingleFieldName(method.Params))
		methodResult = append(methodResult, single)
	}
	return mockTypeStruct + strings.Join(methodResult, "\n")
}

func (this *Generator) generateSingleInit(structInfo StructInfo) string {
	return fmt.Sprintf("RegisterProxyMock([] %s { & %s{}})",
		this.getInterfaceName(structInfo.Name),
		this.getMockName(structInfo.Name))
}
