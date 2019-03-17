package models

import (
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model

	ID          string
	Name        string
	AccessToken string
}
