package repository

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID          uint64       `json:"ID" gorm:"primaryKey"`
	Title       string       `json:"Title" gorm:"not null"`
	Description string       `json:"Description"`
	Actors      []Actor      `json:"Actors" gorm:"many2many:movie_actors"`
	Evaluation  []Evaluation `json:"Evaluation" gorm:"foreignKey:MovieFk"`
}

type MovieRepository struct {
	dbConnection *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		dbConnection: db,
	}
}

func (movieRepo *MovieRepository) GetAll() ([]Movie, error) {
	var res []Movie
	tx := movieRepo.dbConnection.Find(&res)
	if tx.Error != nil {
		return res, tx.Error
	}
	return res, nil
}

func (movieRepo *MovieRepository) FindOne(id uint64) (Movie, error) {
	var res Movie
	tx := movieRepo.dbConnection.First(&res, id)

	if tx.Error != nil {
		return res, tx.Error
	}
	return res, nil

}

func (movieRepo *MovieRepository) Create(elem Movie) (Movie, error) {
	var movieErr Movie

	tx := movieRepo.dbConnection.Create(&elem)
	if tx.Error != nil {
		return movieErr, tx.Error
	}
	return elem, nil
}

func (movieRepo *MovieRepository) Update(movie Movie) error {
	tx := movieRepo.dbConnection.Save(&movie)
	return tx.Error
}

func (movieRepo *MovieRepository) DeleteOne(id uint64) error {
	tx := movieRepo.dbConnection.Unscoped().Delete(&Movie{}, id)
	return tx.Error
}
