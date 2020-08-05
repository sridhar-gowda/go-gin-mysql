package service

import (
	"github.com/sridhar-gowda/go-gin-mysql/models"
	"github.com/sridhar-gowda/go-gin-mysql/repository"
)

const (
	ALIEN = "Alien_user"
)

func CreateCustomer(customer *models.Customer) (err error) {
	if err = repository.Db.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCustomer(customer *models.Customer) (err error) {
	if err = repository.Db.Save(customer).Error; err != nil {
		return err
	}
	return nil
}

func createAlienCustomer(id uint64) (createdId uint64, err error) {
	alien := &models.Customer{
		Id:     id,
		Name:   ALIEN,
		Active: false,
	}
	if err := repository.Db.Create(alien).Error; err != nil {
		return id, err
	}
	return alien.Id, nil
}

func validateCustomer(customerId uint64) (isValidDbEntry bool, active bool) {
	customer := &models.Customer{}
	if err := repository.Db.Where("id =?", customerId).Find(&customer).Error; err == nil {
		return true, customer.Active
	}
	return false, false
}
