package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rozy97/tinder-apps/user-service/config"
)

func main() {
	env := config.InitEnvironment()

	r := gin.Default()

	r.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "register success",
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"access_token": gin.H{
				"aud": "http://api.example.com",
				"iss": "https://krakend.io",
				"sub": "1234567890qwertyuio",
				"jti": "mnb23vcsrt756yuiomnbvcx98ertyuiop",
				"roles": []string{"role_a", "role_b"},
				"exp": time.Now().Add(2 * time.Hour),
			},
			"refresh_token": gin.H{
				"aud": "http://api.example.com",
				"iss": "https://krakend.io",
				"sub": "1234567890qwertyuio",
				"jti": "mnb23vcsrt756yuiomn12876bvcx98ertyuiop",
				"exp": time.Now().Add(7 * 24 * time.Hour),
			},
			"exp": time.Now().Add(2 * time.Hour),
		})
	})

	err := r.Run(":" + env.AppPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
