package models

type Customer struct {
	Id     uint64 `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Name   string `gorm:"type:varchar(255);not null" json:"name"`
	Active bool   `gorm:"type:BOOLEAN;not null" json:"active"`
}

//func (c *Customer) TableName() string {
//	return "customers"
//}
