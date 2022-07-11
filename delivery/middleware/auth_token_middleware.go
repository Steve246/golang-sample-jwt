package middleware

import (
	"fmt"
	"golang-sample-jwt/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	Authorization string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	acctToken utils.Token
}

func (a *authTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H {
				"message": "Unauthorized",
			})
			c.Abort()
		}

		tokenString := strings.Replace(h.Authorization, "Bearer ", "", -1)
		fmt.Println("token String: ", tokenString)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		token, err := a.acctToken.VerifyAccessToken(tokenString)

		userId, err := a.acctToken.FetchAccessToken(token)

		if userId == "" || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return 
		}

		fmt.Println("token: ", token)

		if token != nil {
			c.Set("user-id", userId)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
		}
		
	}
}


func NewTokenValidator(acctToken utils.Token) AuthTokenMiddleware {
	return &authTokenMiddleware{acctToken: acctToken}

}