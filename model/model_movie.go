package model

import (
	"errors"
	"gorm.io/gorm"
)

// Movie the movie model
type Movie []struct {
	MovieID   string `gorm:"primary_key" json:"movie_id"`
	MovieName string `json:"movie_name"`
	Genre     string `json:"genre"`
	Sypnosis  string `json:"sypnosis"`
}

// TableName for gorm
func (Movie) TableName() string {
	return "movies"
}

func (m *Movie) GetAllMovies() error {
	err := DB().Table(m.TableName()).Find(m).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

func (m *Movie) Search(searchString string) error {
	searchSQL := "%" + searchString + "%"
	err := DB().Table(m.TableName()).Where("movie_name LIKE ? OR genre LIKE ? OR Sypnosis LIKE ?", searchSQL, searchSQL, searchSQL).Find(m).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

func (m *Movie) GetMovieByID(id []string) error {
	err := DB().Table(m.TableName()).Where("movie_id IN ?", id).Find(m).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}
