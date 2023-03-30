package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "sql"
	password = "root"
	dbPort   = "3306"
	dbname   = "finalassignment"
	db       *gorm.DB
	err      error
)

func StartDB() {
	db, err := sql.Open("mysql", "username:123456@tcp(127.0.0.1:3306)/finalassigment")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

}

// func StartDB() {
// 	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbPort)
// 	dsn := config
// 	db, err := sql.Open("mysql", "user:password@/dbname")

// 	if err != nil {
// 		panic(err)
// 	}
// }

func GetDB() *gorm.DB {
	return db
}
