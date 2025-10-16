package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(dsn string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := gormDB.DB()
	if err == nil {
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(25)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
	return gormDB, nil
}
