package database

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(config *viper.Viper) *gorm.DB {
	username := config.GetString("DATABASE_USERNAME")
	password := config.GetString("DATABASE_PASSWORD")
	host := config.GetString("DATABASE_HOST")
	port := config.GetInt("DATABASE_PORT")
	database := config.GetString("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db
}

// migrate -database "mysql://root:@tcp(localhost:3306)/url_database_api" -path db/migrations up
// migrate -database "mysql://root:@tcp(localhost:3306)/url_database_api" -path db/migrations down
// migrate create -ext sql -dir db/migrations create_table_todos
// migrate create -ext sql -dir db/migrations create_table_addresses
// migrate create -ext sql -dir db/migrations create_table_users
