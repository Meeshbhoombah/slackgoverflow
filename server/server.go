package main

import (
        "os"
        "log"
        "fmt"
        "time"
        "context"
        "os/signal"

        "github.com/labstack/echo"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/database"
	"github.com/archproj/slackoverflow/slack"
        "github.com/archproj/slackoverflow/routes"
)

const (
	VERSION = "0.1.0"
)

func main() {
	cfg, err := config.Load() // from environment
	if err != nil {
		log.Panic(err)
	}

        e := echo.New()

	db, err := database.Init(cfg)
	if err != nil {
		log.Panic(err)
	}

	sc, err := slack.Init(cfg)
	if err != nil {
		log.Panic(err)
	}

        //routes.Init(cfg, e, db, sc)

	go func() {
                // TODO: add Host, Port, and struct support to config
		if err := e.Start(fmt.Sprintf("%s:%s", "0.0.0.0", "8080")); err != nil {
			e.Logger.Info("shutting down the server.")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
