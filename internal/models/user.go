package models

import (
	"errors"

	"ticket/internal/util/struc"

	"golang.org/x/crypto/bcrypt"
)

func (m *Model) NewUser(user *struc.User) error {
	// first user admin by default
	var count int
	m.db.Model(&struc.User{}).Count(&count)
	if count == 0 {
		user.Admin = true
	}

	//encrypt password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)

	return m.db.Create(user).Error
}

func (m *Model) GetUserByEmailPass(email string, pass string) (*struc.User, error) {
	u := &struc.User{}

	m.db.Where("email = ?", email).First(u)

	//check password
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		return &struc.User{}, errors.New("Password incorrect")
	}

	return u, nil
}

func (m *Model) CountUsers() int {
	var count int
	m.db.Model(&struc.User{}).Count(&count)
	return count
}
