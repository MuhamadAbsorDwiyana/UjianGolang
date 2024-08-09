package main

import (
	db "UjianGolang/Configs"
	models "UjianGolang/Models"
	web "UjianGolang/Routes"
)

func main() {
	// Init database
	db.InitDB()
	db.DB.AutoMigrate(&models.Visitor{}, &models.User{}) // Init every changes structure table

	// Setup fiber routing
	web.Setup()
}
