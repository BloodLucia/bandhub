package controller

import "github.com/gin-gonic/gin"

type userController struct {
}

// Login implements UserController.
func (*userController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

// Register implements UserController.
func (*userController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

type UserController interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
}

func NewUserController() UserController {
	return &userController{}
}
