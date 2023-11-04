package user_info

import (
	"scaler/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct{}

type Reply struct {
	UserId     int    `json:"userId"`
	Username   string `json:"username"`
	Status     string `json:"status"`
	Level      int    `json:"level"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
