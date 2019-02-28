package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/archproj/slackoverflow/config"
	m "github.com/archproj/slackoverflow/middlewares"
	"github.com/archproj/slackoverflow/routes"
	"github.com/archproj/slackoverflow/slack"
)

const (
	VERSION = "0.2.0"
)

func main() {
	// config variables are injected in the environment
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Unable to load environment variables: ", err)
	}

	sc, err := slack.Init(cfg)
	if err != nil {
		log.Fatal("Unable to integrate slackoverflow: ", err)
	}

	e := echo.New()

	e.Use(m.EmbedInContext(cfg, sc))

	routes.Bind(e)

	go func() {
		err := e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
		if err != nil {
			log.Warning("Shutting down server...", err)
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
		log.Fatal(err)
	}
}
