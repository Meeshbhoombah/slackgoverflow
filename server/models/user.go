package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

        MemberId string `gorm:type:varchar(9);not null`
        Email    string `gorm:type:varchar(255);not null`
	Drops    []Drop
}
