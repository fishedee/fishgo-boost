.PHONY:gen
bench:gen
	go test -bench=".*" -benchmem -v 
gen:
	go build
	rm -rf testdata/testdata_querygen.go
	cd testdata && ../querygen ./...
	cd ../../language/testdata && ../../cmd/querygen/querygen ./...
test:gen
	go test . -v
test2:gen
	cd ../../language && go test . -v

