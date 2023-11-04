package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// User 用户基础表
type User struct {
	BaseModel
	Id         int    `json:"id"`          // 主键
	Username   string `json:"username"`    // 登录账号
	Password   string `json:"password"`    // 登录密码
	Name       string `json:"name"`        // 姓名
	Status     string `json:"status"`      // 用户状态
	Level      int    `json:"level"`       // 用户级别
	Avatar     string `json:"avatar"`      // 头像
	Email      string `json:"email"`       // 邮箱
	Phone      string `json:"phone"`       // 手机号
	CreateTime int64  `json:"create_time"` // 创建时间 时间戳：秒
	UpdateTime int64  `json:"update_time"` // 更新时间 时间戳：秒
	DeleteTime int64  `json:"delete_time"` // 删除时间 时间戳：秒
}

func (*User) TableName() string {
	return "user"
}

const (
	UserLevelGeneral = 0
	UserLevelAdmin   = 3
	UserLevelSystem  = 9
)

const (
	UserStatusNormal  = "NORMAL"
	UserStatusDisable = "DISABLED"
)

func (m *User) InitSystemUser() {
	tx := m.Db().Model(m).Scopes(m.NotDeleted).Where("level=?", UserLevelSystem)
	if errors.Is(tx.First(m).Error, gorm.ErrRecordNotFound) {
		ts := time.Now().Unix()
		m.Username = "admin"
		m.Password = "73c2ed591baa6cc7e68ba951138b81113f"
		m.Status = UserStatusNormal
		m.Level = UserLevelSystem
		m.CreateTime = ts
		m.UpdateTime = ts
		m.Db().Create(m)
	}
}

func (m *User) First(query any, args ...any) error {
	return m.Db().Model(m).Scopes(m.NotDeleted).Where(query, args...).First(m).Error
}
