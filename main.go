package main

import (
	"golang-sample-jwt/config"
	"golang-sample-jwt/delivery/middleware"
	"golang-sample-jwt/model"
	"golang-sample-jwt/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}



func main() {
	routerEngine := gin.Default()

	// routerEngine.Use(AuthTokenMiddleware())

	cfg := config.NewConfig()

	tokenService := utils.NewTokenService(cfg.TokenConfig)

	routerGroup := routerEngine.Group("/api")

	routerGroup.POST("/auth/login", func(c *gin.Context){

		var user model.Credential


		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"message": "cant't bind struct",
			})
			return 
		}

		if user.Username == "enigma" && user.Password == "123" {
			token, err := tokenService.CreateAccessToken(&user)

			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return 
			}
			c.JSON(200, gin.H {
				"token": token,
			})
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	})

	protectedGroup := routerGroup.Group("/master", middleware.NewTokenValidator(tokenService).RequireToken())
	protectedGroup.GET("/customer", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "customer",
		})
	})

	protectedGroup.GET("/product", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "product",
		})
	})

	


	err := routerEngine.Run(":8888")

	if err != nil {
		panic(err)
	}



}

