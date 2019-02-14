package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/archproj/slackoverflow/config"
)

func Init(cfg *config.Variables) (*gorm.DB, error) {
	conn := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Dbhost,
		cfg.Dbport,
		cfg.Dbuser,
		cfg.Dbpass,
		cfg.Dbname,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
