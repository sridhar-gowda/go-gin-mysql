package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sridhar-gowda/go-gin-mysql/models"
	"github.com/sridhar-gowda/go-gin-mysql/service"
	"net/http"
)

func DeleteIPBlackList(c *gin.Context) {
	ip := c.Request.URL.Query().Get("ip")
	ipBl := &models.IPBlackList{IP: ip}
	err := service.DeleteBlackListIP(ipBl)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "Deleted Successfully"})
	}

}

func AddIPBlackList(c *gin.Context) {
	var ip models.IPBlackList
	c.BindJSON(&ip)
	err := service.AddBlackListIP(&ip)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, ip)
	}
}

func AddBlackListUser(c *gin.Context) {
	var user models.UserBlackList
	c.BindJSON(&user)
	err := service.AddBlackListUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
