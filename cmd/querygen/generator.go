package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"sort"

	. "github.com/fishedee/fishgo-boost/language"
)

type Generator struct {
	workingDir      string
	initPackageName string
	initPackagePath string
}

type GeneratePackage struct {
	importPackage map[string]bool
	funcName      string
	funcBody      string
	initBody      string
}

func NewGenerator(workingDir string, initPackageName string, initPackagePath string) *Generator {
	result := &Generator{
		workingDir:      workingDir,
		initPackageName: initPackageName,
		initPackagePath: initPackagePath,
	}

	return result
}

func (this *Generator) Run(packages []GeneratePackage) {
	fileName := filepath.Base(this.workingDir) + "_querygen.go"
	filePath := filepath.Join(this.workingDir, fileName)

	//处理导入包
	importPackageMap := map[string]bool{}
	for _, singlePackage := range packages {
		for singleImport, _ := range singlePackage.importPackage {
			importPackageMap[singleImport] = true
		}
	}
	importPackageMap["github.com/fishedee/fishgo-boost/language"] = true
	delete(importPackageMap, this.initPackagePath)
	importPackageList := []string{}
	for singlePackage, _ := range importPackageMap {
		importPackageList = append(importPackageList, "\""+singlePackage+"\"")
	}
	sort.Slice(importPackageList, func(i int, j int) bool {
		return importPackageList[i] < importPackageList[j]
	})
	importBody := Implode(importPackageList, "\n")

	//处理funcBody和initBody
	sort.Slice(packages, func(i int, j int) bool {
		return packages[i].funcName < packages[j].funcName
	})
	var funcBody bytes.Buffer
	var initBody bytes.Buffer
	for _, singlePackage := range packages {
		funcBody.WriteString(singlePackage.funcBody)
		initBody.WriteString(singlePackage.initBody)
	}

	//写入数据
	result := `package ` + this.initPackageName + "\n" +
		"import (\n" + importBody + ")\n" +
		funcBody.String() + "\n" +
		"func init(){\n" +
		initBody.String() + "\n" +
		"}\n"
	oldData, _ := os.ReadFile(filePath)
	if string(oldData) == result {
		return
	}
	err := os.WriteFile(filePath, this.formatSource(result), 0644)
	if err != nil {
		panic(err)
	}
}

func (this *Generator) formatSource(data string) []byte {
	result, err := format.Source([]byte(data))
	if err != nil {
		Throw(1, "format source fail!%v,%v", err, data)
	}
	return result
}
