package repository

import (
	"auth/model"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) FindOrCreate(id, email, name, img string) (*model.User, error) {
	var user model.User

	err := r.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user = model.User{
				ID: id,
				Email: email,
				Name:  name,
				Img:   img,
			}

			err = r.DB.Create(&user).Error
			return &user, err
		}

		return nil, err
	}

	user.Name = name
	user.Img = img
	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
