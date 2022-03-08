package model

import (
	"time"
)

type Answer struct {
	ID         string    `json:"id"`
	Answer     string    `json:"answer"`
	AnswererID string    `json:"answered_id"`
	Answerer   Answerer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AnsweredAt time.Time `json:"answered_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type Answerer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
