package model

type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   int    `json:"status"` // 0: inactive, 1: active, 2: deleted
}

func (u *User) TableName() string {
	return "users"
}
