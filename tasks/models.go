package tasks

import "time"

type Task interface {
	TableName() string
}

type Goal struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Completed bool      `json:"completed"`
}

type Quest struct {
	Goal
	Priority int       `json:"priority"`
	Deadline time.Time `json:"deadline"`
}

type Daily struct {
	Goal
	Priority       int       `json:"priority"`
	NextOccurrence time.Time `json:"next_occurrence"`
	Days           int       `json:"days"`
	FreqType       int       `json:"freq_type"`
	BeforeTime     time.Time `json:"before_time"`
	AfterTime      time.Time `json:"after_time"`
}

// Implement the interface for each task type
func (g Goal) TableName() string  { return "goals" }
func (q Quest) TableName() string { return "quests" }
func (d Daily) TableName() string { return "dailies" }
