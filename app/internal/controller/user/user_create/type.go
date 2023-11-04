package user_create

import (
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
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Reply struct{}
