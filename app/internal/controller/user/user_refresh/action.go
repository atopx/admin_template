package user_refresh

import (
	"app/common/ecode"
	"app/common/logger"
	"app/common/public"
	"app/common/utils"
	"app/internal/model"
	"errors"
	"time"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

func (c *Controller) Deal() (any, error) {
	params := c.Params.(*Params)
	if params.RefreshToken == public.EmptyStr {
		return nil, ecode.NewParamError("无效的Token")
	}

	userToken := new(model.UserToken)
	err := userToken.First("refresh_token=?", params.RefreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.NewParamError("无效的Token")
		}
		return nil, ecode.NewSystemError("系统错误")
	}

	current := time.Now().Local()

	// 重新签发token
	expires := current.Add(24 * time.Hour)
	userToken.IssuedTime = current.Unix()
	userToken.ExpireTime = expires.Unix()
	userToken.AccessToken = utils.SignToken(current, expires, userToken.ExpireTime)
	// 使用AccessToken的过期时间加密
	userToken.RefreshToken = utils.SignToken(current, current.Add(7*24*time.Hour), userToken.ExpireTime)

	// save token
	if err = userToken.Db().Save(userToken).Error; err != nil {
		logger.Error(c.Context(), "save token error", zap.Error(err))
		return nil, ecode.NewSystemError("系统错误")
	}
	return &Reply{
		UserId:       userToken.UserId,
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
		ExpireTime:   userToken.ExpireTime,
	}, nil
}
