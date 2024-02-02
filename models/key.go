package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Key struct {
	gorm.Model

	ID  string `gorm:"primaryKey"`
	Key string

	OrganizationID *string
	ProjectID      *string
	UserID         *string
	Attributes     interface{} `sql:"type:jsonb; not null;" gorm:"type:jsonb; default:'{}'; not null"`

	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}

// Note: Gorm will fail if the function signature
//  does not include `*gorm.DB` and `error`

func (key *Key) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	key.ID = uuid.NewString()
	return
}
