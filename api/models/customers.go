package models

import (
	"errors"
	"unicode"

	"github.com/asaskevich/govalidator"
)

// Customer type is the structure of the users that we will store in the database
type Customer struct {
	ID             string `gorm:"primary_key"`
	Name           string `gorm:"size:255"`
	Email          string `gorm:"size:255; unique"`
	HashedPassword string
}

// CustomerForm is the struct used to login or register
type CustomerForm struct {
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:255"`
}

// ValidateCustomer takes a user as parameter and check if its properties are valid
func ValidateCustomer(customerForm *CustomerForm) error {
	_, err := govalidator.ValidateStruct(customerForm)
	if err != nil {
		return err
	}
	if customerForm.Name == "" {
		return errors.New("Username is empty")
	}

	if valid := verifyPassword(customerForm.Password); valid == false {
		return errors.New("The password must contain at least 8 characters, a capital letter and a number")
	}

	return nil
}

func verifyPassword(s string) bool {
	number := false
	upper := false

	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		default:
		}
	}

	return number && upper && len(s) >= 8
}
