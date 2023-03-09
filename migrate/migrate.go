package main

import "github.com/BrandenM-PM/go-rest-api/initializers"
import "github.com/BrandenM-PM/go-rest-api/models"

func init() {
    initializers.LoadEnvVars()
    initializers.ConnectToDB()
}

func main() {
    initializers.DB.AutoMigrate(&models.Article{})
}
