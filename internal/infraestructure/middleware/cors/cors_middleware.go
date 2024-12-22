package cors

import (
	"github.com/gin-gonic/gin"
	"quotes-api/internal/infraestructure/client/firestore"
	"quotes-api/internal/infraestructure/middleware"
	"quotes-api/internal/util/constant"
	"quotes-api/internal/util/env"
)

type corsMiddleware struct{}

func (t corsMiddleware) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", getCorsOrigin())
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func NewCorsMiddleware() middleware.Middleware {
	return corsMiddleware{}
}

func getCorsOrigin() string {
	var corsOriginValue string
	var err error

	switch env.IsProductionEnv() {
	case true:
		corsOriginValue, err = firestore.GetValue(constant.CorsOriginProd)
	case false:
		corsOriginValue, err = firestore.GetValue(constant.CorsOriginDev)
	}
	if err != nil {
		return ""
	}

	return corsOriginValue
}
