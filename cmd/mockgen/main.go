package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func runProfile(handler func()) {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("start cpu prof")

	// 开始 CPU profile
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// 要剖析的任务
	handler()
}

func mainTask(config Config) {
	parser := NewParser()
	analyseResult := parser.Run(config.packageName)

	generator := NewGenerator(config.fileRegex, config.typeRegex)
	err := generator.Run(analyseResult)
	if err != nil {
		fmt.Println("generate dir error " + err.Error())
		os.Exit(1)
		return
	}
}

func main() {
	config := ReadConfig()
	if config.isProfile {
		runProfile(func() {
			mainTask(config)
		})
	} else {
		mainTask(config)
	}
}
