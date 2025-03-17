package main

import (
	"url-shortener/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3003"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:shortcode", handlers.RedirectURL)
	r.GET("/urls", handlers.GetAllURLs)
	r.POST("/:shortcode/track", handlers.TrackClick)

	r.Run(":8080")
}