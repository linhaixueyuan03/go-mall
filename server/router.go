package server

import (
	"github.com/gin-gonic/gin"
	"go-mall/api"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 用户注册
		v1.POST("user/register", api.UserRegister)
	}
	return r
}
