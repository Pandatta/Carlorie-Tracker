package main

import (
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	// router.Use(gin.logger())
	// router.Use(cors.Default())

	// router.POST("/entry/create", routes.AddEntry)
}
