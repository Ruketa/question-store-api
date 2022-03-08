package model

import (
	"time"
)

type Question struct {
	ID            string     `json:"id"`
	Question      string     `json:"question"`
	QuestionerID  string     `json:"questionner_id"`
	Questioner    Questioner `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	QuestionAt    time.Time  `json:"question_at"`
	Category      string     `json:"category"`
	QuestionTitle string     `json:"question_title"`
	CreatedAt     time.Time  `json:"created_at"`
}

type Questioner struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type QuestionAnswer struct {
	ID         string    `json:"id"`
	AnswerID   string    `json:"answer_id"`
	Answer     Answer    `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	QuestionID string    `json:"question_id"`
	Question   Question  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProjectQuestion struct {
	QuestionID string    `json:"question_id"`
	Question   Question  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ProjectID  string    `json:"project_id"`
	Project    Project   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt  time.Time `json:"created_at"`
}
