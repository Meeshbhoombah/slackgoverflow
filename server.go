package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/archproj/slackoverflow/config"
	"github.com/archproj/slackoverflow/database"
	m "github.com/archproj/slackoverflow/middlewares"
	"github.com/archproj/slackoverflow/routes"
)

const (
	VERSION = "0.2.0"
)

func main() {
	cfg, err := config.Load() // from environment
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	db, err := database.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}

	e.Use(m.EmbedInContext(cfg, db))

	if cfg.Env == `PRODUCTION` {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://slackoverflowmake.herokuapp.com/"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}))
	} else if cfg.Env == `DEVELOPMENT` {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000", "https://https://5c942a8c.ngrok.io"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}))
	}

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
