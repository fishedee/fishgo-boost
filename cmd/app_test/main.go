package main

import (
	"time"

	. "github.com/fishedee/fishgo-boost/app/ioc"
	. "github.com/fishedee/fishgo-boost/app/log"
	. "github.com/fishedee/fishgo-boost/app/workgroup"
)

//go:generate mock ^./model/.*/.*(ao|db)\.go$ ^.*(Ao|Db)$
func main() {
	MustInvokeIoc(func(log Log, server *Server) {
		workgroup, err := NewWorkGroup(log, WorkGroupConfig{
			CloseTimeout: time.Second * 7,
			GraceClose:   true,
		})
		if err != nil {
			panic(err)
		}
		workgroup.Add(server)
		err = workgroup.Run()
		if err != nil {
			panic(err)
		}
	})
}
