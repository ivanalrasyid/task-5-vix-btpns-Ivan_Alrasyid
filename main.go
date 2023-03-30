package main

import (
	"rakamin-final-assigment/database"
	"rakamin-final-assigment/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
