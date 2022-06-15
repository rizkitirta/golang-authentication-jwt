package main

import (
	"golang-authentication-jwt/config"
	"golang-authentication-jwt/controller"
	_"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var(
	db *gorm.DB = config.SetUpDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main()  {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	authRoute := r.Group("api/auth")
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)

	r.Run() 
}