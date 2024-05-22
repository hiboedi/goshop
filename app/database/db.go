package database

import (
	"fmt"
	"log"

	"github.com/hiboedi/zenshop/app/helpers"
	"github.com/hiboedi/zenshop/app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbInit() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	DBName := helpers.GetEnv("DATABASE_NAME", "exam")
	DBPort := helpers.GetEnv("DATABASE_PORT", "5432")
	DBUser := helpers.GetEnv("DATABASE_USER", "boedi")
	DBPassword := helpers.GetEnv("DATABASE_PASSWORD", "")
	DBHost := helpers.GetEnv("DATABASE_HOST", "localhost")
	DBDriver := helpers.GetEnv("DATABASE_Driver", "postgres")

	if DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
		helpers.PanicIfError(err)
		return db
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DBHost, DBUser, DBPassword, DBName, DBPort)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

		if err != nil {
			panic("Failed on connecting to the database server")
		}
		return db
	}
}

func Migrate() {
	db := DbInit()
	err := db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Order{},
	)
	helpers.PanicIfError(err)

	fmt.Println("Migration success")
}
