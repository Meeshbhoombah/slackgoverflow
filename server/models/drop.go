package models

import (
	"github.com/jinzhu/gorm"
)

type Drop struct {
	gorm.Model

        ContentHash string `gorm:type:varchar(256)`
}
