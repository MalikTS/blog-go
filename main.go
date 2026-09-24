package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"blog-go/config"
	"blog-go/handlers"
	"blog-go/middleware"
	"blog-go/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	config.InitDB()
	config.DB.AutoMigrate(&models.User{}, &models.Post{})

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/posts", handlers.GetPosts)
	r.GET("/posts/:id", handlers.GetPost)

	protected := r.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.POST("/posts", handlers.CreatePost)
		protected.PUT("/posts/:id", handlers.UpdatePost)
		protected.DELETE("/posts/:id", handlers.DeletePost)
		protected.GET("/profile", handlers.GetProfile)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}