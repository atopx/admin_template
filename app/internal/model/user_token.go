package model

// UserToken 用户令牌表
type UserToken struct {
	BaseModel
	Id           int    `json:"id"`            // 主键
	UserId       int    `json:"user_id"`       // 用户id
	AccessToken  string `json:"access_token"`  // 登录token
	RefreshToken string `json:"refresh_token"` // 登录token
	IssuedTime   int64  `json:"issued_time"`   // 签发时间 时间戳：秒
	ExpireTime   int64  `json:"expire_time"`   // 过期时间 时间戳：秒
	CreateTime   int64  `json:"create_time"`   // 创建时间 时间戳：秒
	UpdateTime   int64  `json:"update_time"`   // 更新时间 时间戳：秒
	DeleteTime   int64  `json:"delete_time"`   // 删除时间 时间戳：秒
}

func (*UserToken) TableName() string {
	return "user_token"
}

func (m *UserToken) First(query any, args ...any) error {
	return m.Db().Model(m).Scopes(m.NotDeleted).Where(query, args...).First(m).Error
}
