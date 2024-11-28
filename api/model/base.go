package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `gorm:"type:timestamp;" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;" json:"updatedAt"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}

func AllModels() []interface{} {
	return []interface{}{
		&User{},
	}
}

type Querier interface{}
