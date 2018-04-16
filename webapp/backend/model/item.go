package model

import (
	"time"
)

type Item struct {
	ID        int64     `json:"id"`
	CatID     int       `json:"cat_id"`
	LangKey   string    `json:"lang_key"`
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
