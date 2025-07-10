package main

import (
	"barang-bekas/config"
	"barang-bekas/models"
	"barang-bekas/routes"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Product{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
