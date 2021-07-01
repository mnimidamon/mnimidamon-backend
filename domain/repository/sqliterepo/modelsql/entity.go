package modelsql

import (
	"time"
)

type Entity struct {
	ID uint `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
