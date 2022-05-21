package helper

import (
	"context"
	"fmt"
	"time"

	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnectionPostgre(ctx context.Context, data *models.Users) *gorm.DB {

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		data.Host, data.Username, data.Password, data.Dbname, data.Port, "disable")

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

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

	return db

}
