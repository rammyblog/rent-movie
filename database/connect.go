package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rammyblog/rent-movie/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(seed *bool) (*gorm.DB, error) {

	user, exist := os.LookupEnv("DB_USER")

	if !exist {
		log.Fatal("DB_USER not set in .env")
		return nil, fmt.Errorf("DB_USER not set in .env")

	}

	pass, exist := os.LookupEnv("DB_PASSWORD")

	if !exist {
		log.Fatal("DB_PASSWORD not set in .env")
		return nil, fmt.Errorf("DB_PASSWORD not set in .env")

	}

	port, exist := os.LookupEnv("DB_PORT")

	if !exist {
		log.Fatal("DB_PASS not set in .env")
		return nil, fmt.Errorf("DB_PORT not set in .env")

	}

	host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		log.Fatal("DB_HOST not set in .env")
		return nil, fmt.Errorf("DB_HOST not set in .env")

	}

	name, exist := os.LookupEnv("DB_NAME")

	if !exist {
		log.Fatal("DB_NAME not set in .env")
		return nil, fmt.Errorf("DB_NAME not set in .env")

	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, pass, name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	DB = db
	MigrateTables(db)
	if *seed {
		seedMovies(db)
	}
	return db, nil

}

func MigrateTables(db *gorm.DB) {
	fmt.Println("Migrating tables")
	// Auto migrate tables here

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
		panic(err)
	}
	if err := db.AutoMigrate(&models.Genre{}); err != nil {
		log.Fatal(err)
		panic(err)
	}
	if err := db.AutoMigrate(&models.Movie{}); err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err := db.AutoMigrate(&models.Rent{}); err != nil {
		log.Fatal(err)
		panic(err)
	}

}

func seedMovies(db *gorm.DB) {
	fmt.Println("Seeding data")
	defer fmt.Println("Done Seeding")

	movies := []models.Movie{
		{
			Name: "Game Night",
			Year: 2018,
			Genre: []models.Genre{
				{Name: "Drama"},
			},
			Rating:      7.2,
			ReleaseDate: time.Now().UTC(),
			Summary:     "A group of friends who meet regularly for game nights find themselves trying to solve a murder mystery.",
			Image:       "https://images-na.ssl-images-amazon.com/images/M/MV5BMjQxMDE5NDg0NV5BMl5BanBnXkFtZTgwNTA5MDE2NDM@._V1_SY500_CR0,0,337,500_AL_.jpg",
			Available:   true,
		},
		{
			Name: "Area X: Annihilation",
			Year: 2018,
			Genre: []models.Genre{
				{Name: "Adventure"},
			},
			Rating:      5.5,
			Available:   true,
			ReleaseDate: time.Now().Add(5).UTC(),
			Summary:     `A biologist's husband disappears. She puts her name forward for an expedition into an environmental disaster zone, but does not find what she"s expecting. The expedition team is made up of the biologist, an anthropologist, a psychologist, a surveyor, and a linguist.`,
			Image:       "https://images-na.ssl-images-amazon.com/images/M/MV5BMTk2Mjc2NzYxNl5BMl5BanBnXkFtZTgwMTA2OTA1NDM@._V1_SY500_CR0,0,320,500_AL_.jpg",
		},
		{
			Name: "Hannah",
			Year: 2017,
			Genre: []models.Genre{
				{Name: "Fantasy"},
			},
			Rating:      5.5,
			Available:   true,
			ReleaseDate: time.Now().Add(2).UTC(),
			Summary:     `Intimate portrait of a woman drifting between reality and denial when she is left alone to grapple with the consequences of her husband"s imprisonment.`,
			Image:       "https://images-na.ssl-images-amazon.com/images/M/MV5BNWJmMWIxMjQtZTk0Mi00YTE0LTkyNzAtYzQxYjcwYjE4ZDk2XkEyXkFqcGdeQXVyMTc4MzI2NQ@@._V1_SY500_SX350_AL_.jpg",
		},
		{
			Name: "The Lodgers",
			Year: 2017,
			Genre: []models.Genre{
				{Name: "Horror"},
			},
			Rating:      9.4,
			ReleaseDate: time.Now().Add(2).UTC(),
			Summary:     "1920, rural Ireland. Anglo Irish twins Rachel and Edward share a strange existence in their crumbling family estate. Each night, the property becomes the domain of a sinister presence (The Lodgers) which enforces three rules upon the twins: they must be in bed by midnight; they may not permit an outsider past the threshold; if one attempts to escape, the life of the other is placed in jeopardy. When troubled war veteran Sean returns to the nearby village, he is immediately drawn to the mysterious Rachel, who in turn begins to break the rules set out by The Lodgers. The consequences pull Rachel into a deadly confrontation with her brother - and with the curse that haunts them.",

			Image: "https://images-na.ssl-images-amazon.com/images/M/MV5BM2FhM2E1MTktMDYwZi00ODA1LWI0YTYtN2NjZjM3ODFjYmU5XkEyXkFqcGdeQXVyMjY1ODQ3NTA@._V1_SY500_CR0,0,337,500_AL_.jpg",
		},
	}
	for _, movie := range movies {
		db.Create(&movie)
	}
}
