package models

import "time"

type HourlyStat struct {
	Id           uint64    `gorm:"primary_key;AUTO_INCREMENT;not null"`
	CustomerID   uint64    `gorm:"index:customer_idx"`
	RequestCount uint64    `gorm:"default:0;not null"`
	InvalidCount uint64    `gorm:"default:0;not null"`
	Time         time.Time `gorm:"type:timestamp"`
}

type PostRequest struct {
	CustomerID uint64 `json:"customerID" binding:"required"`
	TagID      uint64 `json:"tagID" binding:"required"`
	UserID     string `json:"userID" binding:"required"`
	RemoteIP   string `json:"remoteIP" binding:"required"`
	Timestamp  int64  `json:"timestamp" binding:"required"`
}

type GetStats struct {
	CustomerID string     `json:"customerID" binding:"required"`
	Date       *time.Time `json:"date" binding:"required"`
}

type Response struct {
	Customer     string
	HourlyCounts []HourlyCount
}

type HourlyCount struct {
	RequestCount uint64
	InvalidCount uint64
	Time         time.Time
}
