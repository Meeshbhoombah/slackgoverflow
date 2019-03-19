package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/archproj/slackoverflow/config"
)

func Init(cfg *config.Variables) (*gorm.DB, error) {
	conn := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s dbname=%s password=%s",
		cfg.Dbhost,
		cfg.Dbport,
		cfg.Dbuser,
		cfg.Dbname,
		cfg.Dbpass,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Workspace{})

	return db, nil
}
