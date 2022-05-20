package databases

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	err        error
	connection string
)

func Initpostgre(ctx context.Context) *gorm.DB {

	if db == nil {
		connection = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", "127.0.0.1", "postgres", "123", "bulk", 2345, "disable")

		db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

		if err != nil {
			panic("Failed to connect database postgre")
		}

		sqlDB, err := db.DB()

		if err != nil {
			panic("Failed to create pool connection database postgre")
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return db

}
