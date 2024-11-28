package model

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(255);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Email    string `gorm:"type:varchar(255);" json:"email"`
	Status   int    `gorm:"type:int;not null;default:1" json:"status"` // 0: inactive, 1: active, 2: deleted
}

func (u *User) TableName() string {
	return "users"
}
