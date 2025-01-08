package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rajithacharith/go-template/internal/pkg/app"
	"github.com/sirupsen/logrus"
)

func main() {
	app, err := app.MakeAPIServer()
	if err != nil {
		logrus.Fatalf("Failed to create app: %v", err)
	}

	go func() {
		logrus.WithField("port", 3000).Info("Starting Hello World Service...")
		err := app.Listen(fmt.Sprintf(":%d", 3000))
		if err != nil {
			logrus.Fatalf("Failed to start server: %v", err)
		}
	}()

	sigtermC := make(chan os.Signal, 1)
	signal.Notify(sigtermC, os.Interrupt, syscall.SIGTERM)

	<-sigtermC
	logrus.Info("SIGTERM received: gracefully shutting down...")

	if err := app.Shutdown(); err != nil {
		logrus.Errorf("server shutdown error: %v", err)
	}
}
