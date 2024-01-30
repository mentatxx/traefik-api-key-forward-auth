package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Key struct {
	gorm.Model

	Id  string `gorm:"primaryKey"`
	Key string
}

// Note: Gorm will fail if the function signature
//  does not include `*gorm.DB` and `error`

func (key *Key) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	key.Id = uuid.NewString()
	return
}
