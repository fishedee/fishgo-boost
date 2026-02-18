package main

import (
	"os"
)

type Config struct {
	packageName string
	fileRegex   string
	typeRegex   string
	isProfile   bool
}

func ReadConfig() Config {
	argv := os.Args
	argv = argv[1:]
	if len(argv) < 3 {
		panic("mockgen [packageName] [fileRegex] [typeRegex] {-profile}")
	}
	var config Config
	config.packageName = argv[0]
	config.fileRegex = argv[1]
	config.typeRegex = argv[2]
	if len(argv) >= 4 && argv[3] == "-profile" {
		config.isProfile = true
	} else {
		config.isProfile = false
	}
	return config
}
