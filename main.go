package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sridhar-gowda/go-gin-mysql/config"
	"github.com/sridhar-gowda/go-gin-mysql/repository"
	"github.com/sridhar-gowda/go-gin-mysql/router"
)

var err error

func main() {
	repository.Db, err = gorm.Open("mysql", repository.DbURL(repository.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}
	defer repository.Db.Close()
	repository.HandleMigration()

	servicePort := fmt.Sprintf(":%d", config.GetConf.ServicePort)
	r := router.SetupRouter()
	//running
	r.Run(servicePort)
}
