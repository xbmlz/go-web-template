package model

type Role struct {
	BaseModel
	Name   string `gorm:"unique;not null;size:32" json:"name"`
	Code   string `gorm:"unique;not null;size:32" json:"code"`
	Desc   string `gorm:"size:32" json:"desc"`
	Status int8   `gorm:"type:int;not null;default:1" json:"status"` // 0: inactive, 1: active, 2: deleted
}

func (Role) TableName() string {
	return "sys_role"
}

type UserRole struct {
	BaseModel
	UserID uint `gorm:"not null"`
	RoleID uint `gorm:"not null"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}
