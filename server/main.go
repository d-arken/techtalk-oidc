package main

import (
	"encoding/gob"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"techtalk-oidc/auth"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}


func setupRouter() *gin.Engine {
	
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Use(auth.JWTProtected()).GET("/authenticated-ping", func(c *gin.Context) {
		claims := c.MustGet(auth.OIDCClaimsContext).(auth.OIDCClaims)
		c.JSON(200, gin.H{"claims": claims})
	})
	return r
}

func main() {
	err := godotenv.Load()
	gob.Register(oidc.IDToken{})
	if err != nil {
		panic(err.Error())
	}

	r := setupRouter()
	r.Run(":8080")
}