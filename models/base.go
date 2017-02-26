package models

import "time"

type BaseModel struct {
	ID        int64     `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
