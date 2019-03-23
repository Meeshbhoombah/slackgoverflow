package models

import (
	"github.com/jinzhu/gorm"
)

type Workspace struct {
	gorm.Model

	TeamID    string
	TeamName  string
	UserToken string
	ChanName  string
	ChanID    string
}
