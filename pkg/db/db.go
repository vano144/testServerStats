package db

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/logrusadapter"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	config "testServerStats/configs"
	"time"
)

var Db *sqlx.DB

// pgxLogLevel convert log level to pgx log level
func pgxLogLevel(ll log.Level) pgx.LogLevel {
	switch ll {
	case log.DebugLevel:
		return pgx.LogLevelDebug
	case log.InfoLevel:
		return pgx.LogLevelInfo
	case log.WarnLevel:
		return pgx.LogLevelWarn
	case log.ErrorLevel:
		return pgx.LogLevelError
	case log.TraceLevel:
		return pgx.LogLevelTrace
	default:
		return pgx.LogLevelNone
	}
}

func openSqlxConnPool() (*sqlx.DB, error) {
	ll, _ := log.ParseLevel(config.Config.LogLevel)

	connConfig := pgx.ConnConfig{
		Host:     config.Config.DatabaseHost,
		Port:     uint16(config.Config.DatabasePort),
		Database: config.Config.DatabaseName,
		User:     config.Config.DatabaseUser,
		Password: config.Config.DatabasePass,

		Logger:   logrusadapter.NewLogger(log.StandardLogger()),
		LogLevel: pgxLogLevel(ll),

		PreferSimpleProtocol: false,
	}

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: config.Config.MaxConnections, // max simultaneous connections to use
		AcquireTimeout: time.Duration(config.Config.AcquireTimeout) * time.Second,

		// custom callback for dealing after connection has been established
		AfterConnect: nil,
	})

	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDBFromPool(connPool)
	db.SetMaxIdleConns(config.Config.MaxIdleConnections) // set the maximum idle conns size
	db.SetMaxOpenConns(config.Config.MaxOpenConnections) // set the maximum size of the pool

	return sqlx.NewDb(db, "pgx"), nil
}

func InitDb() *sqlx.DB {
	var err error

	Db, err = openSqlxConnPool()
	if err != nil {
		log.WithError(err).Fatal("Failed to establish database connection")
	}

	return Db
}
