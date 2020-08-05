package service

import (
	"github.com/sridhar-gowda/go-gin-mysql/models"
	"github.com/sridhar-gowda/go-gin-mysql/repository"
)

func AddBlackListIP(ip *models.IPBlackList) (err error) {
	if err = repository.Db.Create(ip).Error; err != nil {
		return err
	}
	return nil
}

func AddBlackListUser(user *models.UserBlackList) (err error) {
	if err = repository.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBlackListIP(ip *models.IPBlackList) error {
	if err := getBlackListIP(ip); err != nil {
		return err
	}
	if err1 := repository.Db.Delete(ip).Error; err1 != nil {
		return err1
	}
	return nil
}

func getBlackListIP(ip *models.IPBlackList) error {
	err1 := repository.Db.Where("ip =?", ip.IP).Find(&ip).Error
	return err1
}

func isValidIP(ip string) bool {
	ipBl := &models.IPBlackList{}
	if err := repository.Db.Where("ip =?", ip).Find(&ipBl).Error; err == nil {
		return false
	}
	return true
}

func isValidUserAgent(user string) bool {
	userAgent := &models.UserBlackList{}
	if err := repository.Db.Where("name =?", user).Find(&userAgent).Error; err == nil {
		return false
	}
	return true
}
