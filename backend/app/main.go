package app

import (
	"context"

	"rustdesk-api-server-pro/app/middleware"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/errcode"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

func newApp(cfg *config.ServerConfig, dbEngine *xorm.Engine) (*iris.Application, error) {
	app := iris.Default()
	if dbEngine == nil {
		return nil, errcode.New(errcode.ERRB001.Code, errcode.ERRB001.Message)
	}
	app.RegisterDependency(dbEngine, cfg)

	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		context.Application().Logger().Infof("(404)▶ %s:%s", context.Method(), context.Request().URL.Path)
	})

	app.Use(iris.Compression)
	app.Use(middleware.ContainerLogRecorder(app))
	if cfg.HttpConfig.PrintRequestLog {
		app.Use(middleware.RequestLogger(cfg.DebugMode))
	}

	SetRoute(app)

	app.HandleDir("/", iris.Dir(cfg.HttpConfig.StaticDir))

	return app, nil
}

func StartServer() (bool, error) {
	return StartServerWithContext(context.Background())
}

func StartServerWithContext(ctx context.Context) (bool, error) {
	cfg := config.GetServerConfig()
	if config.IsUnsafeSignKey(cfg.SignKey) {
		return false, errcode.New(errcode.ERRB002.Code, errcode.ERRB002.Message)
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

	middleware.RecordContainerEvent(dbEngine, "INFO", "system", "Server started on port "+cfg.HttpConfig.Port)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				_ = r
			}
		}()
		<-ctx.Done()
		_ = app.Shutdown(ctx)
	}()

	err = app.Listen(cfg.HttpConfig.Port, iris.WithoutBodyConsumptionOnUnmarshal)
	if err != nil && ctx.Err() != nil {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
