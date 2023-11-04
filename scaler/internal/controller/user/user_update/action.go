package user_update

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
	user := new(model.User)
	if err := user.First("id=?", c.UserId()); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.NewParamError("用户不存在")
		}
		return nil, ecode.NewSystemError("系统错误")
	}

	if user.Level >= model.UserLevelSystem && user.Id != c.UserId() {
		return nil, ecode.NewParamError("禁止修改其他超级管理员")
	}

	user.Level = params.Level
	user.Phone = params.Phone
	user.Email = params.Phone
	user.Avatar = params.Avatar
	user.Password = utils.Hash(params.Password)
	user.UpdateTime = time.Now().Unix()
	user.Db().Model(user).Updates(user)
	return &Reply{}, nil
}
