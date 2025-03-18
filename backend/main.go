package main

import (
	"http"
	"url-shortener/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests.",
        },
        []string{"method", "endpoint", "status"},
    )
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3003"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	http.Handle("/metrics", promhttp.Handler())

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:shortcode", handlers.RedirectURL)
	r.GET("/urls", handlers.GetAllURLs)
	r.POST("/:shortcode/track", handlers.TrackClick)

	r.Run(":8080")
}