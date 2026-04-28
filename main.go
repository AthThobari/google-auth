package main

import (
	"auth/config"
	"auth/handler"
	"auth/model"
	"auth/pb"
	"auth/repository"
	"auth/service"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// ======================
	// CONFIG
	// ======================
	cfg := config.Load()

	// ======================
	// DB
	// ======================
	db := config.InitDB(os.Getenv("DB"))
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("failed migrate:", err)
	}

	// ======================
	// REPOSITORY
	// ======================
	userRepo := repository.NewUserRepository(db)

	// ======================
	// SERVICE
	// ======================
	authService := &service.AuthService{
		ClientID: cfg.GoogleClientID,
		UserRepo: userRepo,
	}

	// ======================
	// HANDLER
	// ======================
	authHandler := &handler.AuthHandler{
		Service:   authService,
		JWTSecret: cfg.JWTSecret,
	}

	// ======================
	// GIN ROUTER
	// ======================
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"POST", "GET", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/login/google", func(c *gin.Context) {
		var req struct {
			IdToken string `json:"id_token"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "bad request"})
			return
		}

		resp, err := authHandler.LoginWithGoogle(c, &pb.GoogleLoginRequest{
			IdToken: req.IdToken,
		})

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, resp)
	})

	r.Run(":8080")
}