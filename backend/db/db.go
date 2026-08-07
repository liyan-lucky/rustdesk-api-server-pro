package db

import (
	"rustdesk-api-server-pro/config"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

var DbEngine *xorm.Engine

func NewEngine(cfg *config.DbConfig) (*xorm.Engine, error) {
	dsn := cfg.Dsn
	if cfg.Driver == "sqlite" && !strings.Contains(dsn, "_pragma") {
		sep := "?"
		if strings.Contains(dsn, "?") {
			sep = "&"
		}
		dsn = dsn + sep + "_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)"
	}
	engine, err := xorm.NewEngine(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}
	location, _ := time.LoadLocation(cfg.TimeZone)
	engine.TZLocation = location
	engine.DatabaseTZ = location
	engine.ShowSQL(cfg.ShowSql)
	if cfg.Driver == "sqlite" {
		engine.SetMaxIdleConns(1)
		engine.SetMaxOpenConns(1)
	} else {
		engine.SetMaxIdleConns(100)
		engine.SetMaxOpenConns(100)
	}
	DbEngine = engine
	return engine, nil
}
