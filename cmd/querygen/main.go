package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	. "github.com/fishedee/fishgo-boost/language"
)

var (
	recursive = flag.Bool("r", false, "generate package including sub package")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of querygen:\n")
	fmt.Fprintf(os.Stderr, "\tlarge improve performance language/query.go function \n")
	fmt.Fprintf(os.Stderr, "\tquerygen [packageName]\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttps://github.com/fishedee/fishgo-boost\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("querygen fail: ")
	defer CatchCrash(func(e Exception) {
		log.Fatal(e.Error())
	})

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
		panic("need package name")
	}

	//parse
	genPackages := []GeneratePackage{}
	parser := NewParser()
	parser.Run(args[0], func(callInfo *CallInfo) {
		funcFullName := callInfo.PackagePath + "." + callInfo.FuncName
		genPackage := handleQueryGen(funcFullName, callInfo)
		if genPackage != nil {
			genPackages = append(genPackages, *genPackage)
		}
	})

	//generate
	generator := NewGenerator(parser.workingDir, parser.initPackageName, parser.initPackagePath)
	generator.Run(genPackages)
}
