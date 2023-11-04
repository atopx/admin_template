package user_disable

import (
	"app/common/ecode"
	"app/internal/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (c *Controller) Deal() (any, error) {
	user := new(model.User)
	if err := user.First("id=?", c.UserId()); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.NewParamError("用户不存在")
		}
		return nil, ecode.NewSystemError("系统错误")
	}
	if user.Level >= model.UserLevelSystem {
		return nil, ecode.NewParamError("禁止禁用超级管理员")
	}
	user.Db().Model(user).Updates(map[string]any{
		"status":      model.UserStatusDisable,
		"update_time": time.Now().Unix(),
	})
	return &Reply{}, nil
}
