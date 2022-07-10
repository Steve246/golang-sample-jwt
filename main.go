package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type Credential struct {
	Username string 	`json:"username"`
	Password string   `json:"password"`
}

func main() {
	routerEngine := gin.Default()

	routerEngine.Use(AuthTokenMiddleware())

	routerGroup := routerEngine.Group("/api")

	routerGroup.POST("/auth/login", func(c *gin.Context){

		var user Credential

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"message": "cant't bind struct",
			})
			return 
		}

		if user.Username == "enigma" && user.Password == "123" {
			c.JSON(200, gin.H {

				"token": "123456",

			})
		} else {
			c.AbortWithStatus(401)
		}

	})

	routerGroup.GET("/customer", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "customer",
		})
	})

	routerGroup.GET("/product", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "product",
		})
	})

	


	err := routerEngine.Run(":8888")

	if err != nil {
		panic(err)
	}



}


func AuthTokenMiddleware () gin.HandlerFunc {

	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/auth/login" {
			c.Next()
			fmt.Println("sss")
		} else {
			h := AuthHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H {
					"message": "Unauthorized",
				})
				c.Abort()
			}
			if h.AuthorizationHeader == "123456" {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "token Invalid",
				})
				c.Abort()

			}
		}
	}

}