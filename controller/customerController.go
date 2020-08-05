package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sridhar-gowda/go-gin-mysql/models"
	"github.com/sridhar-gowda/go-gin-mysql/service"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var customer models.Customer
	jsonErr := c.BindJSON(&customer)
	if customer.Name == "" || jsonErr != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("customer name empty"))
		return
	}
	err := service.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, customer)
	}

}

func UpdateCustomerStatus(c *gin.Context) {
	var customer models.Customer
	c.BindJSON(&customer)
	Id, strErr := strconv.ParseUint(c.Param("id"), 0, 64)
	if customer.Name == "" || strErr != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("wrong input"))
		return
	}
	customer.Id = Id
	err := service.UpdateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
