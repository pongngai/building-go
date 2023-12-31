package _const

const (
	DatabaseGeneratedPath = "infrastructure/db.go"
	Header                = `// Code generated by go generate; DO NOT EDIT.
// This file was generated by running: go generate`
	DatabaseTemplate = `
package infrastructure

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Database Database
}

type Database struct {
	Host          string
	Port          uint64
	Username      string
	Password      string
	DBName        string
	SSLMode       string
	TimeZone      string
	Schema        string
	GormLogEnable bool
	MaxIdleConn   int
	MaxOpenConn   int
}

var appConfig Config

func NewDB() *gorm.DB {

	err := viper.Unmarshal(&appConfig)

	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Infof("db config %v", appConfig)
		panic(err)
	}

	log.Debugf("loaded database configuration = %v", appConfig)
	var gormLogger logger.Interface

	if !appConfig.Database.GormLogEnable {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else {
		gormLogger = logger.Default
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v search_path=%v", appConfig.Database.Host, appConfig.Database.Username, appConfig.Database.Password, appConfig.Database.DBName, appConfig.Database.Port, appConfig.Database.SSLMode, appConfig.Database.TimeZone, appConfig.Database.Schema)), &gorm.Config{
		Logger: gormLogger,
	})
	sqlDB, _ := db.DB()
	maxIdleConn, maxOpenConn := defaultConnectionPool(appConfig)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func defaultConnectionPool(cfg Config) (int, int) {
	var maxIdleConn int
	var maxOpenConn int
	if cfg.Database.MaxIdleConn == 0 {
		maxIdleConn = 10
	}

	if cfg.Database.MaxOpenConn == 0 {
		maxOpenConn = 10
	}

	return maxIdleConn, maxOpenConn
}`
)
