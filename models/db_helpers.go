package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const host = "localhost"
const port = 5432
const user = "postgres"
const dbname = "seymour"
const password = ""

func GetDB() *gorm.DB {
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func RunMigrations() {
	db := GetDB()

	defer db.Close()

	db.AutoMigrate(
		&FeedUrl{},
		&FeedPull{},
		&User{},
		&Profile{},
		&FeedItem{},
	)
}

func MigrateToZero() {
	db := GetDB()
	db.DropTable(
		&FeedUrl{},
		&FeedPull{},
		&User{},
		&Profile{},
		&FeedItem{},
	)
}
