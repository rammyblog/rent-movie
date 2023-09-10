package worker

import (
	"log"
	"time"

	"github.com/rammyblog/rent-movie/database"
	"github.com/rammyblog/rent-movie/models"
)

func SendEmailReminder() {

	ticker := time.NewTicker(24 * time.Hour)
	log.Println("Worker started")
	for ; ; <-ticker.C {
		var rents []models.Rent

		if err := database.DB.Preload("User").Preload("Movie").Where("returned = ? AND due_date < ?", false, time.Now().Format(time.RFC3339)).Find(&rents).Error; err != nil {
			log.Fatal("Error occurred while fetching movies")
		}

		for _, rent := range rents {
			// send mail to user
			// Might just as well spawn another go routine
			log.Printf("%v is yet to return %v with movie id of %v and rent id of %v", rent.User.Name, rent.Movie.Name, rent.MovieId, rent.ID)
		}

	}
}
