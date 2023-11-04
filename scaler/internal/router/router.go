package router

import (
	"net/http"
	"scaler/common/interceptor"
	v1 "scaler/internal/api/v1"

	"github.com/gin-gonic/gin"
)

func Route(app *gin.Engine) {

	v1Group := app.Group("/api/v1")
	{
		// 开放接口
		v1Group.POST("/user/login", v1.UserLogin)

		// 通用接口
		generalGroup := v1Group.Use(interceptor.Auth(1))
		{
			generalGroup.GET("/user/info", v1.UserInfo)
			generalGroup.POST("/user/refresh", v1.UserRefresh)
		}

		// 管理员接口
		adminGroup := v1Group.Use(interceptor.Auth(3))
		{
			adminGroup.POST("/user/list", v1.UserList)
		}

		// 超管接口
		systemGroup := v1Group.Use(interceptor.Auth(9))
		{
			systemGroup.POST("/user/create", v1.UserCreate)
			systemGroup.POST("/user/update", v1.UserUpdate)
			systemGroup.DELETE("/user/delete", v1.UserDelete)
			systemGroup.PATCH("/user/disable", v1.UserDisable)
		}
	}

	// 外部接口
	openGroup := app.Group("/openapi")
	{
		openGroup.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
	}
}
