package user_list

import (
	"app/common/public"
	"app/internal/model"
	"fmt"
	"time"
)

func (c *Controller) Deal() (any, error) {
	fmt.Println(111)
	params := c.Params.(*Params)
	user := new(model.User)
	tx := user.Db().Model(user).Scopes(user.NotDeleted)

	if params.Filter.Keyword != public.EmptyStr {
		key := user.Like(params.Filter.Keyword)
		tx.Where("username like ? or email like ? or phone like ?", key, key, key)
	}

	if params.Filter.Level > 0 {
		tx.Where("level = ?", params.Filter.Level)
	}

	if params.Filter.TimeRange.Right > 0 {
		tx.Where("update_time between ? and ?", params.Filter.TimeRange.Left, params.Filter.TimeRange.Right)
	}

	tx.Count(&params.PageInfo.Count)

	if params.PageInfo.Size == 0 {
		params.PageInfo.Size = 15
	}

	if params.PageInfo.Index == 0 {
		params.PageInfo.Index = 1
	}
	users := make([]model.User, 0)
	tx.Scopes(user.Paginate(&params.PageInfo)).Find(&users)

	reply := &Reply{
		PageInfo: params.PageInfo,
		List:     make([]record, 0, len(users)),
	}
	for _, item := range users {
		reply.List = append(reply.List, record{
			UserId:     item.Id,
			Username:   item.Username,
			Status:     item.Status,
			Level:      item.Level,
			Avatar:     item.Avatar,
			Email:      item.Email,
			Phone:      item.Phone,
			CreateTime: time.Unix(item.CreateTime, 0).Format(time.DateTime),
			UpdateTime: time.Unix(item.UpdateTime, 0).Format(time.DateTime),
		})
	}
	return reply, nil
}
