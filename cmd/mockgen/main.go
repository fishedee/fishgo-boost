package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func mainTask() {

	data, err := ReadDir(".")
	if err != nil {
		fmt.Println("read dir error " + err.Error())
		os.Exit(1)
		return
	}

	data, err = FilterDir(data)
	if err != nil {
		fmt.Println("filter dir error " + err.Error())
		os.Exit(1)
		return
	}

	err = Generator(data)
	if err != nil {
		fmt.Println("generate dir error " + err.Error())
		os.Exit(1)
		return
	}
}

func runProfile() {
	if Config.isProfile {
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
		mainTask()
	} else {
		mainTask()
	}
}
func main() {
	err := ReadConfig()
	if err != nil {
		fmt.Println("read config error " + err.Error())
		os.Exit(1)
		return
	}
	runProfile()
}
