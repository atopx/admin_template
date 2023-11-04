package v1

import (
	"app/internal/controller"
	"app/internal/controller/user/user_create"
	"app/internal/controller/user/user_delete"
	"app/internal/controller/user/user_disable"
	"app/internal/controller/user/user_info"
	"app/internal/controller/user/user_list"
	"app/internal/controller/user/user_login"
	"app/internal/controller/user/user_refresh"
	"app/internal/controller/user/user_update"

	"github.com/gin-gonic/gin"
)

// UserLogin
// @summary 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body user_login.Params true "请求参数"
// @Response 200 object public.Response{data=user_login.Reply} "调用成功"
// @Router /api/v1/user/login [post]
func UserLogin(ctx *gin.Context) {
	controller.Handle(user_login.NewController(ctx))
}

// UserRefresh
// @summary 刷新token
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body user_refresh.Params true "请求参数"
// @Response 200 object public.Response{data=user_refresh.Reply} "调用成功"
// @Router /api/v1/user/refresh [post]
func UserRefresh(ctx *gin.Context) {
	controller.Handle(user_refresh.NewController(ctx))
}

// UserInfo
// @summary 用户详情
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body user_info.Params true "请求参数"
// @Response 200 object public.Response{data=user_info.Reply} "调用成功"
// @Router /api/v1/user/info [get]
func UserInfo(ctx *gin.Context) {
	controller.Handle(user_info.NewController(ctx))
}

// UserList
// @summary 用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body user_list.Params true "请求参数"
// @Response 200 object public.Response{data=user_list.Reply} "调用成功"
// @Router /api/v1/user/list [post]
func UserList(ctx *gin.Context) {
	controller.Handle(user_list.NewController(ctx))
}

// UserDelete
// @summary 删除用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body user_delete.Params true "请求参数"
// @Response 200 object public.Response{data=user_delete.Reply} "调用成功"
// @Router /api/v1/user/delete [delete]
func UserDelete(ctx *gin.Context) {
	controller.Handle(user_delete.NewController(ctx))
}

// UserUpdate
// @summary 更新用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body user_update.Params true "请求参数"
// @Response 200 object public.Response{data=user_update.Reply} "调用成功"
// @Router /api/v1/user/update [post]
func UserUpdate(ctx *gin.Context) {
	controller.Handle(user_update.NewController(ctx))
}

// UserCreate
// @summary 创建用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body user_create.Params true "请求参数"
// @Response 200 object public.Response{data=user_create.Reply} "调用成功"
// @Router /api/v1/user/create [post]
func UserCreate(ctx *gin.Context) {
	controller.Handle(user_create.NewController(ctx))
}

// UserDisable
// @summary 禁用用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body user_disable.Params true "请求参数"
// @Response 200 object public.Response{data=user_disable.Reply} "调用成功"
// @Router /api/v1/user/disable [patch]
func UserDisable(ctx *gin.Context) {
	controller.Handle(user_disable.NewController(ctx))
}
