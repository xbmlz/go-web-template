package model

import (
	"github.com/xbmlz/go-web-template/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"size:32;not null" json:"username"`
	Password string `gorm:"size:128;not null" json:"password"`
	Email    string `gorm:"size:32;" json:"email"`
	Status   int8   `gorm:"type:int;not null;default:1" json:"status"` // 0: inactive, 1: active, 2: deleted
	Roles    []Role `gorm:"many2many:sys_user_role;" json:"roles"`
}

func (User) TableName() string {
	return "sys_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = utils.HashPassword(u.Password)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password = utils.HashPassword(u.Password)
	}
	return
}
