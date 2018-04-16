package model

import (
	"time"
)

type Item struct {
	ID        int       `json:"id"`
	CatID     int       `json:"cat_id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Details   string    `json:"details"`
	WebURL    string    `json:"web_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (m *Item) TableName() string {
	return "item"
}
