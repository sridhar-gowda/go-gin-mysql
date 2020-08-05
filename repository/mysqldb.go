package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sridhar-gowda/go-gin-mysql/config"
	"github.com/sridhar-gowda/go-gin-mysql/models"
)

var Db *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     config.GetConf.DbHost,
		Port:     config.GetConf.DbPort,
		User:     config.GetConf.DbUsername,
		Password: config.GetConf.DbPassword,
		DBName:   config.GetConf.DbDbname,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func HandleMigration() {
	//repository.Db.DropTableIfExists(&models.HourlyStat{}, &models.Customer{})
	Db.AutoMigrate(&models.Customer{}, &models.HourlyStat{}, &models.IPBlackList{}, &models.UserBlackList{})
	Db.Model(&models.HourlyStat{}).AddForeignKey("customer_id", "customers(`id`)", "CASCADE", "RESTRICT")
	Db.Exec("ALTER TABLE hourly_stats ADD CONSTRAINT unique_customer_time UNIQUE (customer_id, time);")
}
