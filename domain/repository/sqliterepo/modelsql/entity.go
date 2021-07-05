package modelsql

import (
	"time"
)

type Entity struct {
	ID uint `gorm:"primaryKey; uniqueIndex"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// Uncomment for Soft delete support.
	// DeletedAt gorm.DeletedAt
}
