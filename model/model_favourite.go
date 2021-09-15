package model

import (
	"errors"
	"gorm.io/gorm"
)

// Movie the movie model
type Favourite []struct {
	FavID   int   `gorm:"primary_key" json:"fav_id"`
	UserID  string `gorm:"foreign_key" json:"user_id"`
	MovieID string `json:"movie_id"`
}

// TableName for gorm
func (Favourite) TableName() string {
	return "favourites"
}

func (f *Favourite) GetFavourites(UserID int) error{
	err := DB().Table(f.TableName()).Where("user_id=?", UserID).Find(f).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

func (f *Favourite) SaveFavourite() error {
	db := DB().Table(f.TableName()).Create(f)

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return ErrKeyConflict
	}

	return nil
}
