package main

import (
	"context"
	"net/http"

	. "github.com/fishedee/fishgo-boost/app/config"
	. "github.com/fishedee/fishgo-boost/app/gzip"
	. "github.com/fishedee/fishgo-boost/app/ioc"
	. "github.com/fishedee/fishgo-boost/app/log"
	. "github.com/fishedee/fishgo-boost/app/middleware"
	. "github.com/fishedee/fishgo-boost/app/render"
	. "github.com/fishedee/fishgo-boost/app/router"
	. "github.com/fishedee/fishgo-boost/app/session"
	. "github.com/fishedee/fishgo-boost/app/validator"
	"github.com/fishedee/fishgo-boost/cmd/tool/util"
	. "github.com/fishedee/fishgo-boost/language"
)

type Server struct {
	log           Log
	config        ServerConfig
	server        *http.Server
	routerFactory *RouterFactory
}

type ServerConfig struct {
	Listen string `config:"listen"`
}

func NewServer(
	trigger util.Trigger,
	config Config,
	log Log,
	validatorFactory ValidatorFactory,
	sessionFactory SessionFactory,
	renderFactory RenderFactory) *Server {

	serverConfig := ServerConfig{}
	config.MustBind("", &serverConfig)

	routerFactory := NewRouterFactory()
	routerFactory.Use(NewPProfMiddleware())
	routerFactory.Use(NewLogMiddleware(log, nil))

	if config.MustString("runmode") == "dev" {
		gzip, err := NewGzip(GzipConfig{})
		if err != nil {
			panic(err)
		}
		routerFactory.Use(NewGzipMiddleware(gzip))
		QueryReflectWarning(false)
	} else {
		//FIXME 暂时不打开优化功能
		QueryReflectWarning(false)
	}

	routerFactory.GET("/hello", func(w http.ResponseWriter, r *http.Request, param RouterParam) {
		w.Write([]byte("hello world!"))
	})

	return &Server{
		log:           log,
		config:        serverConfig,
		routerFactory: routerFactory,
	}
}

func (this *Server) Run() error {
	this.log.Debug("server is running... listen %v", this.config.Listen)
	this.server = &http.Server{
		Addr:    this.config.Listen,
		Handler: this.routerFactory.Create(),
	}
	return this.server.ListenAndServe()
}

func (this *Server) Close() {
	this.server.Shutdown(context.TODO())
}

func init() {
	MustRegisterIoc(NewServer)
}
