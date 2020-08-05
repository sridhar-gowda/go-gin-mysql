package models

type IPBlackList struct {
	IP string `gorm:"primary_key;type:varchar(255)" json:"ip"`
}

type UserBlackList struct {
	Name string `gorm:"primary_key;type:varchar(255)" json:"user-agent"`
}
