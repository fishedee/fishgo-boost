.PHONY:gen
bench:gen
	go test . -bench=. -benchmem -v 
gen:
	go install github.com/fishedee/fishgo-boost/language/querygen
	rm -rf testdata/testdata_querygen.go
	querygen -r github.com/fishedee/fishgo-boost/language/querygen/testdata
test:gen
	go test . -v
