package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sridhar-gowda/go-gin-mysql/helper"
	"github.com/sridhar-gowda/go-gin-mysql/models"
	"github.com/sridhar-gowda/go-gin-mysql/service"
	"net/http"
)

func PostRequest(c *gin.Context) {
	var request models.PostRequest
	jsonErr := c.BindJSON(&request)
	err := service.PostRequest(&request, jsonErr)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"success": request})
	}
}

func GetStats(c *gin.Context) {
	customerID := c.Request.URL.Query().Get("customerID")
	date := helper.StringToDate(c.Request.URL.Query().Get("date"))

	if customerID == "" || date == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("wrong json"))
	}

	request := &models.GetStats{
		CustomerID: customerID,
		Date:       date,
	}

	stats, err := service.GetDayStats(request)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, stats)
	}
}
