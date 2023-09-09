package models

type Genre struct {
	Model
	Name string `gorm:"not null;uniqueIndex;" json:"name"`
}


// monolight django app - handling admissions 2016() -refactor Change things
//  requirements that changed over the year
// rebuild
// 3-4 months
// 
