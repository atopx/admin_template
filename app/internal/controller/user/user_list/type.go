package user_list

import (
	"app/common/public"
	"app/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct {
	PageInfo public.PageInfo `json:"pageInfo"`
	Filter   struct {
		Keyword   string           `json:"keyword"`
		Status    string           `json:"status"`
		Level     int              `json:"level"`
		TimeRange public.TimeRange `json:"timeRange"`
	} `json:"filter"`
}

type Reply struct {
	PageInfo public.PageInfo `json:"pageInfo"`
	List     []record        `json:"list"`
}

type record struct {
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
