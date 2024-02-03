package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"quotes-api/internal/infraestructure/middleware/cors"
)

func StartApp() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	router := gin.Default()
	router.Use(cors.NewCorsMiddleware().Execute())
	mapURLs(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return router.Run(":" + port)
}
