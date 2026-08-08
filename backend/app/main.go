package app

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"rustdesk-api-server-pro/app/middleware"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/util"

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

	if cfg.HttpConfig.TLSPort != "" {
		go startTLSServer(app, cfg)
	}

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

// startTLSServer 在 goroutine 中启动 HTTPS 反向代理服务器，用于 WebAuthn 等需要安全上下文的功能
func startTLSServer(app *iris.Application, cfg *config.ServerConfig) {
	certFile := cfg.HttpConfig.TLSCertFile
	keyFile := cfg.HttpConfig.TLSKeyFile

	if certFile == "" || keyFile == "" {
		certFile = filepath.Join(filepath.Dir(cfg.Db.Dsn), "tls-cert.pem")
		keyFile = filepath.Join(filepath.Dir(cfg.Db.Dsn), "tls-key.pem")
	}

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		hosts := cfg.HttpConfig.TLSHosts
		if len(hosts) == 0 {
			hosts = []string{"localhost", "127.0.0.1"}
		}
		app.Logger().Infof("生成自签名证书: hosts=%v", hosts)
		if _, _, err := util.GenerateSelfSignedCert(certFile, keyFile, hosts); err != nil {
			app.Logger().Errorf("生成自签名证书失败: %v", err)
			return
		}
	}

	httpPort := strings.TrimPrefix(cfg.HttpConfig.Port, ":")
	if httpPort == "" {
		httpPort = "8080"
	}
	target := &url.URL{Scheme: "http", Host: "127.0.0.1:" + httpPort}
	proxy := httputil.NewSingleHostReverseProxy(target)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Set("X-Forwarded-Proto", "https")
	}

	tlsServer := &http.Server{
		Addr:    cfg.HttpConfig.TLSPort,
		Handler: proxy,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
		ReadHeaderTimeout: 10 * time.Second,
	}

	app.Logger().Infof("HTTPS 反向代理监听: %s -> http://127.0.0.1:%s", cfg.HttpConfig.TLSPort, httpPort)
	if err := tlsServer.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
		app.Logger().Errorf("HTTPS 服务器启动失败: %v", err)
	}
}
