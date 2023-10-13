package v1

import "github.com/gin-gonic/gin"

// RegisterGuestAPIRouter 这里的路由不需要登录也能访问
func (ar *APIRouter) RegisterGuestAPIRouter(r *gin.RouterGroup) {
	r.POST("/regiser", ar.userController.Register())
	r.POST("/login", ar.userController.Login())
}
