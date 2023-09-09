package models

import (
	"time"
)

type Movie struct {
	Model
	Name        string    `json:"name"`
	Year        int       `json:"year"`
	Rating      float64   `json:"rating"`
	ReleaseDate time.Time `json:"releaseDate"`
	Summary     string    `json:"summary"`
	Available   bool      `json:"available"`
	Image       string    `json:"image"`
	Genre       []Genre   `gorm:"many2many:movie_genre;" json:"genre"`
}
