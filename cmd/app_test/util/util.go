package util

import (
	"github.com/fishedee/fishgo-boost/app/config"
	"github.com/fishedee/fishgo-boost/app/ioc"
	"github.com/fishedee/fishgo-boost/app/log"
	"github.com/fishedee/fishgo-boost/app/metric"
	"github.com/fishedee/fishgo-boost/app/proxy"
	"github.com/fishedee/fishgo-boost/app/render"
	"github.com/fishedee/fishgo-boost/app/session"
	"github.com/fishedee/fishgo-boost/app/sqlf"
	"github.com/fishedee/fishgo-boost/app/trigger"
	"github.com/fishedee/fishgo-boost/app/validator"
)

type Log = log.Log

type Config = config.Config

type Session = session.Session

type Validator = validator.Validator

type Render = render.Render

type Trigger = trigger.Trigger

func MustRegisterIoc(obj interface{}) {
	ioc.MustRegisterIoc(obj)
}

func MustRegisterIocWithProxy(obj interface{}) {
	ioc.MustRegisterIoc(proxy.WrapCreatorWithProxy(obj))
}

func NewConfig() Config {
	config, err := config.NewConfig("ini", "data/conf/config.ini")
	if err != nil {
		panic(err)
	}
	return config
}

func NewLog(config Config) Log {
	var logConfig log.LogConfig
	config.MustBind("log", &logConfig)
	log, err := log.NewLog(logConfig)
	if err != nil {
		panic(err)
	}
	return log
}

func NewSessionFactory(config Config) session.SessionFactory {
	var jwtTokenConfig session.JwtTokenConfig
	config.MustBind("session", &jwtTokenConfig)
	sessionFactory, err := session.NewJwtTokenFactory(jwtTokenConfig)
	if err != nil {
		panic(err)
	}
	return sessionFactory
}

func NewValidatorFactory(config Config) validator.ValidatorFactory {
	var validatorConfig validator.ValidatorConfig
	config.MustBind("validator", &validatorConfig)
	validatorFactory, err := validator.NewValidatorFactory(validatorConfig)
	if err != nil {
		panic(err)
	}
	return validatorFactory
}

func NewRenderFactory(config Config) render.RenderFactory {
	renderFactory, err := render.NewRenderFactory(render.RenderConfig{})
	if err != nil {
		panic(err)
	}
	return renderFactory
}
func NewTrigger() trigger.Trigger {
	trigger, err := trigger.NewTrigger()
	if err != nil {
		panic(err)
	}
	return trigger
}

func NewSqlf(config Config, log Log, metric metric.Metric) sqlf.SqlfDB {
	var sqlfConfig sqlf.SqlfDBConfig
	config.MustBind("sql", &sqlfConfig)
	sqlf, err := sqlf.NewSqlfDB(log, metric, sqlfConfig)
	if err != nil {
		panic(err)
	}
	return sqlf
}

func init() {
	ioc.MustRegisterIoc(NewConfig)
	ioc.MustRegisterIoc(NewLog)
	ioc.MustRegisterIoc(NewSessionFactory)
	ioc.MustRegisterIoc(NewValidatorFactory)
	ioc.MustRegisterIoc(NewRenderFactory)
	ioc.MustRegisterIoc(NewTrigger)
}
