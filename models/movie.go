package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name        string
	Year        int
	Rating      float64
	ReleaseDate time.Time
	Summary     string
	Available   bool
	Image       string
	Genre       []Genre `gorm:"many2many:movie_genre;"`
}
