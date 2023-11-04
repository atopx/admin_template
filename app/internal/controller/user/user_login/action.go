package user_login

import (
	"app/common/ecode"
	"app/common/public"
	"app/common/utils"
	"app/internal/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (c *Controller) Deal() (any, error) {
	params := c.Params.(*Params)
	if params.Username == public.EmptyStr {
		return nil, ecode.NewParamError("无效的用户名")
	}
	if params.Password == public.EmptyStr {
		return nil, ecode.NewParamError("无效的密码")
	}
	user := new(model.User)
	err := user.First("username=?", params.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.NewParamError("未注册的用户")
		}
		return nil, ecode.NewSystemError("系统错误")
	}
	if utils.Hash(params.Password) != user.Password {
		return nil, ecode.NewParamError("登录密码错误")
	}
	userToken := new(model.UserToken)
	err = userToken.First("user_id=?", user.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ecode.NewSystemError("系统错误")
	}

	current := time.Now().Local()
	ts := current.Unix()

	if userToken.ExpireTime <= ts {
		expires := current.Add(24 * time.Hour)
		userToken.UserId = user.Id
		userToken.IssuedTime = ts
		userToken.ExpireTime = expires.Unix()
		userToken.AccessToken = utils.SignToken(current, expires, userToken.ExpireTime)
		userToken.RefreshToken = utils.SignToken(current, current.Add(7*24*time.Hour), userToken.ExpireTime)
	}

	if err = userToken.Db().Model(userToken).Save(userToken).Error; err != nil {
		return nil, ecode.NewSystemError("系统错误")
	}

	return &Reply{
		UserId:       userToken.UserId,
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
		ExpireTime:   userToken.ExpireTime,
	}, nil
}
