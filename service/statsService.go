package service

import (
	"github.com/sridhar-gowda/go-gin-mysql/helper"
	"github.com/sridhar-gowda/go-gin-mysql/models"
	"github.com/sridhar-gowda/go-gin-mysql/repository"
)

func PostRequest(request *models.PostRequest, jsonErr error) (err error) {

	isTotallyValid, isNewEntry := requestValidation(jsonErr, request)

	if isNewEntry == true {
		createdId, err := createAlienCustomer(request.CustomerID)
		if err != nil {
			return err
		}
		if err1 := postStats(createdId, 0, 1, request.Timestamp); err != nil {
			return err1
		}
	} else {
		var validCount, invalidCount uint64
		if isTotallyValid == true {
			validCount = 1
		} else {
			invalidCount = 1
		}
		stat := getStats(request.CustomerID, request.Timestamp)
		if stat == nil {
			if err := postStats(request.CustomerID, validCount, invalidCount, request.Timestamp); err != nil {
				return err
			}
		} else {
			if err := updateStats(stat.Id, request.CustomerID, stat.RequestCount+validCount, stat.InvalidCount+invalidCount, request.Timestamp); err != nil {
				return err
			}
		}
	}
	return nil
}

func GetDayStats(request *models.GetStats) (stat *models.Response, err error) {

	end := request.Date.AddDate(0, 0, 1)
	var stats []models.HourlyStat
	repository.Db.Where("time >= ? and time < ? and customer_id = ?", request.Date, end, request.CustomerID).Find(&stats)

	var hourlystats []models.HourlyCount

	for _, h := range stats {
		hourlystat := models.HourlyCount{
			RequestCount: h.RequestCount,
			InvalidCount: h.InvalidCount,
			Time:         h.Time,
		}
		hourlystats = append(hourlystats, hourlystat)
	}
	dayStats := &models.Response{
		Customer:     request.CustomerID,
		HourlyCounts: hourlystats,
	}
	return dayStats, nil
}

func requestValidation(jsonErr error, request *models.PostRequest) (isTotallyValid bool, isNewRecord bool) {
	isValidDbEntry, isActive := validateCustomer(request.CustomerID)
	isValidIP := isValidIP(request.RemoteIP)
	isValidUser := isValidUserAgent(request.UserID)

	var isValidJson bool
	if jsonErr != nil {
		isValidJson = false
	} else {
		isValidJson = true
	}

	return isValidDbEntry && isActive && isValidIP && isValidUser && isValidJson, !isValidDbEntry
}

func postStats(customerId uint64, requestCount uint64, invalidCount uint64, timestamp int64) (err error) {

	stats := &models.HourlyStat{
		CustomerID:   customerId,
		RequestCount: requestCount,
		InvalidCount: invalidCount,
		Time:         helper.GetStartHour(timestamp),
	}
	if err := repository.Db.Create(stats).Error; err != nil {
		return err
	}
	return nil
}

func updateStats(Id uint64, customerId uint64, requestCount uint64, invalidCount uint64, timestamp int64) (err error) {
	stats := &models.HourlyStat{
		Id:           Id,
		CustomerID:   customerId,
		RequestCount: requestCount,
		InvalidCount: invalidCount,
		Time:         helper.GetStartHour(timestamp),
	}
	if err := repository.Db.Save(stats).Error; err != nil {
		return err
	}
	return nil
}

func getStats(customerId uint64, timestamp int64) (stat *models.HourlyStat) {
	stats := &models.HourlyStat{}
	start := helper.GetStartHour(timestamp)
	if err := repository.Db.Where("time = ? and customer_id = ?", start, customerId).Find(stats).Error; err != nil {
		return nil
	}
	return stats
}
