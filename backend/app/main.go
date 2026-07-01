package app

import (
	"errors"
	"rustdesk-api-server-pro/app/middleware"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

func newApp(cfg *config.ServerConfig, dbEngine *xorm.Engine) (*iris.Application, error) {
	app := iris.Default()
	if dbEngine == nil {
		return nil, errors.New("db engine is nil")
	}
	app.RegisterDependency(dbEngine, cfg)

	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		context.Application().Logger().Infof("(404)▶ %s:%s", context.Method(), context.Request().URL.Path)
	})

	app.Use(iris.Compression)
	if cfg.HttpConfig.PrintRequestLog {
		app.Use(middleware.RequestLogger(cfg.DebugMode))
	}

	SetRoute(app)

	app.HandleDir("/", iris.Dir(cfg.HttpConfig.StaticDir))

	return app, nil
}

func StartServer() (bool, error) {
	cfg := config.GetServerConfig()
	if config.IsUnsafeSignKey(cfg.SignKey) {
		return false, errors.New("unsafe signKey: set a unique random signKey with at least 32 characters before starting the server")
	}

	dbEngine, err := db.NewEngine(cfg.Db)
	if err != nil {
		return false, err
	}

	app, err := newApp(cfg, dbEngine)
	if err != nil {
		return false, err
	}

	if err := StartJobs(cfg, dbEngine); err != nil {
		return false, err
	}

	err = app.Listen(cfg.HttpConfig.Port, iris.WithoutBodyConsumptionOnUnmarshal)
	if err != nil {
		return false, err
	}

	return true, nil
}
