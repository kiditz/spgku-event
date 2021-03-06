package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

// AddUser used tto register new user
func AddUser(user *e.User) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		if user.Type == "talent" {
			tx.Create(&e.Talent{
				UserID:     user.ID,
				IsVerified: false,
			})
		}
		if user.Type == "company" {
			tx.Create(&e.Company{
				UserID:    user.ID,
				Name:      user.Name,
				IsUpdated: false,
			})
		}
		return nil
	})
}

// FindUserByEmail is used to query user by email address
func FindUserByEmail(email string) (e.User, error) {
	var user e.User
	if db.DB.Where("email = ?", email).Find(&user).RecordNotFound() {
		err := fmt.Errorf("user_not_found")
		return user, err
	}
	return user, nil
}

// FindUserByID is used to query user by email address
func FindUserByID(userID uint) (e.User, error) {
	var user e.User
	if db.DB.Where("id = ?", userID).Find(&user).RecordNotFound() {
		err := fmt.Errorf("user_not_found")
		return user, err
	}
	return user, nil
}
