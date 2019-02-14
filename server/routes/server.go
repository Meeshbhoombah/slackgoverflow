package routes

import (
        "os"
        "log"
        "time"
        "context"
        "os/signal"

        "github.com/labstack/echo"

        "github.com/archproj/slackoverflow/config"
        "github.com/archproj/slackoverflow/slack"
)

func Serve(cfg *config.Variables) {
        e := echo.New()

        sc, err := slack.Init(cfg)
	if err != nil {
		log.Panic(err)
	}

        log.Println(sc)

        go func() {
                if err := e.Start(":8080"); err != nil {
                        e.Logger.Warnf("Shutting down the server with error:%v", err)
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
