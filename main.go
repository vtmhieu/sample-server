package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	PORT          = "8080"
	CLIENT_ORIGIN = "http://localhost:3000"
)

var (
	server *gin.Engine
)

type Config struct {
	ServerPort   string `mapstructure:"PORT"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"published_at"`
}

func main() {
	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))
	router := gin.Default()
	server.Use(cors.New(corsConfig))
	router = server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to VCS Server Management System"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	router.Run()
}
