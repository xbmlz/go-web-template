package model

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp;" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;" json:"updatedAt"`
}

func AllModels() []interface{} {
	return []interface{}{
		&User{},
	}
}

type Querier interface{}
