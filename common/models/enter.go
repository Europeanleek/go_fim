package models

import "time"

type Model struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
}
