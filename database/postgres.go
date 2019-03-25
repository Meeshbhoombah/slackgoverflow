package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/models"
)

func Init(cfg *config.Variables) (*gorm.DB, error) {
	conn := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s dbname=%s password=%s",
		cfg.DB.Dbhost,
		cfg.DB.Dbport,
		cfg.DB.Dbuser,
		cfg.DB.Dbname,
		cfg.DB.Dbpass,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Drop{}, &models.User{})

	return db, nil
}
