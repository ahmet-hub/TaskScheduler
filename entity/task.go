package entity

import "time"

type Task struct {
	Id          string    `json:"Id"`
	Name        string    `json:"Name"`
	StartDate   time.Time `json:"StartDate"`
	EndDate     time.Time `json:"EndDate"`
	WorkingTime int
}
