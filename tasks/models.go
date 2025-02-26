package tasks

import "time"

type Goal struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Completed bool      `json:"completed"`
}

type Quest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Completed bool      `json:"completed"`
	Priority  int       `json:"priority"`
	Deadline  time.Time `json:"deadline"`
}

type Daily struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	Completed      bool      `json:"completed"`
	Priority       int       `json:"priority"`
	NextOccurrence time.Time `json:"next_occurrence"`
	Days           int       `json:"days"`
	FreqType       int       `json:"freq_type"`
	BeforeTime     time.Time `json:"before_time"`
	AfterTime      time.Time `json:"after_time"`
}
