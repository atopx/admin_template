package user_info

import (
	"scaler/common/ecode"
	"scaler/internal/model"
	"time"
)

func (c *Controller) Deal() (any, error) {
	user := new(model.User)
	if err := user.First("id=?", c.UserId()); err != nil {
		return nil, ecode.NewSystemError("系统错误")
	}
	reply := &Reply{
		UserId:     user.Id,
		Username:   user.Username,
		Status:     user.Status,
		Level:      user.Level,
		Avatar:     user.Avatar,
		Email:      user.Email,
		Phone:      user.Phone,
		CreateTime: time.Unix(user.CreateTime, 0).Format(time.DateTime),
		UpdateTime: time.Unix(user.UpdateTime, 0).Format(time.DateTime),
	}
	return reply, nil
}
