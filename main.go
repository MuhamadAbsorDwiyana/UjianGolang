package main

import (
	db "github.com/MuhamadAbsorDwiyana/UjianGolang/Configs"
	models "github.com/MuhamadAbsorDwiyana/UjianGolang/Models"
	web "github.com/MuhamadAbsorDwiyana/UjianGolang/Routes"
)

func main() {
	// Init database
	db.InitDB()
	db.DB.AutoMigrate(&models.Visitor{}, &models.User{}) // Init every changes structure table

	// Setup fiber routing
	web.Setup()
}
