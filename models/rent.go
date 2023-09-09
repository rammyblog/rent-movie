package models

import (
	"time"
)

type Rent struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	DueDate      time.Time `json:"dueDate"`
	Returned     bool      `json:"returned"`
	ReturnedAt   time.Time `json:"returnedAt"`
	DateBorrowed time.Time `json:"dateBorrowed"`
	MovieId      int       `json:"movieId"`
	Movie        Movie     `gorm:"constraint:OnDelete:CASCADE;" json:"movie"`
	UserID       int       ` json:"userId"`
	User         User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
}
