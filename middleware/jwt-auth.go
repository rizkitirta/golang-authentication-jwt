package middleware

import (
	"golang-authentication-jwt/helper"
	"golang-authentication-jwt/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token provided",nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[user_id] : ", claims["user_id"])
			log.Println("Claims[issuer] : ", claims["issuer"])
		}else {
			log.Println(err)
			response := helper.BuildErrorResponse("Invalid token", err.Error(),nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}