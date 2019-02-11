package main

import (
	"github.com/archproj/slackoverflow/config"
	"log"
)

func main() {
	var c config.Variables
	err := c.Load(); if err != nil {
		log.Fatal(err)
	}
}
