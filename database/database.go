package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"microapp-fiber-kit/config"

	"log"
	"os"
	"time"
)

const (
	DefaultMaxOpenConns = 25
	DefaultMaxIdleConns = 25
)

// Database struct
type Database struct {
	gormDb *gorm.DB
}

func (d Database) DB() *gorm.DB {
	return d.gormDb
}

func NewDatabase(cfg *config.Config) *Database {
	newLogger := gLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gLogger.Config{
			SlowThreshold:             time.Second,    // Slow SQL threshold
			LogLevel:                  gLogger.Silent, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound errors for logger
			Colorful:                  true,           // Disable color
		},
	)
	db, err := gorm.Open(dialector(cfg.Database), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()

	pool, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Database.MaxOpen <= 0 {
		cfg.Database.MaxOpen = DefaultMaxOpenConns
	}
	if cfg.Database.MaxIdle <= 0 {
		cfg.Database.MaxIdle = DefaultMaxIdleConns
	}
	pool.SetMaxOpenConns(cfg.Database.MaxOpen)
	pool.SetMaxIdleConns(cfg.Database.MaxIdle)
	pool.SetConnMaxLifetime(5 * time.Minute)

	return &Database{
		gormDb: db,
	}
}

func dialector(cfg config.DatabaseConfig) gorm.Dialector {
	switch cfg.Type {
	case "sqlite":
		return sqlite.Open("sqlite.db")
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database,
		)
		return mysql.New(mysql.Config{
			DriverName:                "mysql",
			DSN:                       dsn,
			SkipInitializeWithVersion: false,
			DefaultStringSize:         255,  // change it if needed
			DisableDatetimePrecision:  true, // true, because datetime precision requires MySQL 5.6
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
		})
	}

	log.Fatalf("'%s' driver is not supported.", cfg.Type)
	return nil
}
