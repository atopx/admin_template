package model

import (
	"fmt"
	"scaler/common/public"

	"gorm.io/gorm"
)

type BaseModel struct{}

func (b *BaseModel) Db() *gorm.DB {
	return public.GetHandler().Db
}

func (dao *BaseModel) NotDeleted(tx *gorm.DB) *gorm.DB {
	return tx.Where("delete_time = 0")
}

func (dao *BaseModel) Like(value string) string {
	return fmt.Sprintf("%%%s%%", value)
}

func (dao *BaseModel) Paginate(pageInfo *public.PageInfo) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageInfo.Disabled {
			return db
		}
		offset := (pageInfo.Index - 1) * pageInfo.Size
		return db.Offset(offset).Limit(pageInfo.Size)
	}
}
