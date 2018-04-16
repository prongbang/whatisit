package model

import "time"

type Category struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (m *Category) TableName() string {
	return "category"
}
