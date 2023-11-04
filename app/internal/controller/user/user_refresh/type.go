package user_refresh

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
	RefreshToken string `json:"refresh_token"`
}

type Reply struct {
	UserId       int    `json:"userId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireTime   int64  `json:"expireTime"`
}
