package main

var (
	queryGenMapper = map[string]queryGenHandler{}
)

type queryGenHandler func(request *CallInfo) *GeneratePackage

func handleQueryGen(name string, request *CallInfo) *GeneratePackage {
	handler, isExist := queryGenMapper[name]
	if isExist == false {
		return nil
	}
	return handler(request)
}

func registerQueryGen(name string, handler queryGenHandler) {
	queryGenMapper[name] = handler
}
