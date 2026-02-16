package main

import (
	"errors"
	"os"
)

var Config struct {
	fileregex string
	typeregex string
	isProfile bool
}

func ReadConfig() error {
	argv := os.Args
	argv = argv[1:]
	if len(argv) < 2 {
		return errors.New("need a file regex argument and a type name regex argument")
	}
	Config.fileregex = argv[0]
	Config.typeregex = argv[1]
	if len(argv) >= 3 && argv[2] == "-profile" {
		Config.isProfile = true
	} else {
		Config.isProfile = false
	}
	return nil
}
