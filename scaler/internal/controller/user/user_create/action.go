package user_create

import (
	"errors"
	"gorm.io/gorm"
	"scaler/common/ecode"
	"scaler/common/utils"
	"scaler/internal/model"
	"time"
)

func (c *Controller) Deal() (any, error) {
	params := c.Params.(*Params)

	if len(params.Username) < 5 {
		return nil, ecode.NewParamError("无效的用户名")
	}

	if len(params.Password) < 5 {
		return nil, ecode.NewParamError("无效的密码")
	}

	if params.Level != model.UserLevelSystem &&
		params.Level != model.UserLevelGeneral &&
		params.Level != model.UserLevelAdmin {
		return nil, ecode.NewParamError("无效的用户级别")
	}

	err := new(model.User).First("username=?", params.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ecode.NewSystemError("系统错误")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ecode.NewParamError("用户名已存在")
	}

	ts := time.Now().Unix()
	user := &model.User{
		Username:   params.Username,
		Password:   utils.Hash(params.Password),
		Name:       params.Name,
		Status:     model.UserStatusNormal,
		Level:      params.Level,
		Avatar:     params.Avatar,
		Email:      params.Email,
		Phone:      params.Phone,
		CreateTime: ts,
		UpdateTime: ts,
	}
	user.Db().Model(user).Create(user)
	return &Reply{}, nil
}
