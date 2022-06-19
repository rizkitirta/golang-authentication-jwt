package controller

import (
	"golang-authentication-jwt/dto"
	"golang-authentication-jwt/entity"
	"golang-authentication-jwt/helper"
	"golang-authentication-jwt/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JWTService
}

func NewAuthController(authService service.AuthService,jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService: jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
		var loginDTO dto.LoginDTO
		if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
			response := helper.BuildErrorResponse("Failed to prosess request", err.Error(),helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		checkCredential := c.authService.VerifyCredentials(loginDTO.Email, loginDTO.Password); 
		if v,ok := checkCredential.(entity.User); ok {
			token, err := c.jwtService.GenerateToken(strconv.Itoa(v.ID))
			
			if err != nil {
				response := helper.BuildErrorResponse("Failed to generate token", err.Error(),helper.EmptyObj{})
				ctx.JSON(http.StatusInternalServerError, response)
				return
			}
			v.Token = token
			response := helper.BuildResponse(true,"Login Success", v)
			ctx.JSON(http.StatusOK, response)
		}else{
			response := helper.BuildErrorResponse("Please check again your credential", "Invalid email or password",helper.EmptyObj{})
			ctx.JSON(http.StatusUnauthorized, response)
		}		 
}

func (c *authController) Register(ctx *gin.Context) {
	 var registerDTO dto.RegisterDTO
	 if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		 response := helper.BuildErrorResponse("Failed to prosess request", err.Error(),helper.EmptyObj{})
		 ctx.JSON(http.StatusBadRequest, response)
		 return
	 }

	 if c.authService.IsDuplicateEmail(registerDTO.Email) {
		 response := helper.BuildErrorResponse("Failed to prosess request", "Email already exist",helper.EmptyObj{})
		 ctx.JSON(http.StatusConflict, response)
		 return
	 }

	 user, err := c.authService.CreateUser(registerDTO);
	 if err != nil {
		 response := helper.BuildErrorResponse("Failed to prosess request", err.Error(),helper.EmptyObj{})
		 ctx.JSON(http.StatusInternalServerError, response)
		 return
	 }

	 token, err := c.jwtService.GenerateToken(strconv.Itoa(user.ID))
	 if err != nil {
		 response := helper.BuildErrorResponse("Failed to generate token", err.Error(),helper.EmptyObj{})
		 ctx.JSON(http.StatusInternalServerError, response)
		 return
	 }

	 user.Token = token
	 response := helper.BuildResponse(true,"Register Success", user)
	 ctx.JSON(http.StatusOK, response)
}