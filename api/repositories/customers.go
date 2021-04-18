package repositories

import (
	"errors"

	"github.com/jasongauvin/zebi-scraper/api/database"
	"github.com/jasongauvin/zebi-scraper/api/models"
	"github.com/jinzhu/gorm"
)

func CreateCustomer(customer *models.Customer) error {
	err := database.DB.Debug().Create(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

// FindCustomerByEmail finds a customer in the db thanks to its email
func FindCustomerByEmail(customer *models.Customer) error {
	err := database.DB.Debug().Save(&customer).Take(&customer).Error
	if err != nil {
		return  err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Customer Not Found")
	}
	return err
}
