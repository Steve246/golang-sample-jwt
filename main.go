package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func main() {
	routerEngine := gin.Default()

	routerGroup := routerEngine.Group("/api")

	routerGroup.GET("/customer", func(ctx *gin.Context){

		authHeader := AuthHeader{}

		if err := ctx.ShouldBindJSON(&authHeader); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return 
		}

		if authHeader.AuthorizationHeader == "123456" {
			ctx.JSON(http.StatusOK, gin.H {
				"message": "customer",
			})
			return 
		}


		ctx.JSON(http.StatusUnauthorized, gin.H {
			"message": "Unauthorized",
		})

	})

	routerGroup.GET("/product", func(ctx *gin.Context) {
		authHeader := AuthHeader{} 

		if err := ctx.ShouldBindJSON(&authHeader); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H {
				"message": "Unauthorized",
			})
			return 
		}
	})


	err := routerEngine.Run(":8888")

	if err != nil {
		panic(err)
	}



}