package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rajithacharith/go-template/internal/pkg/app"
	"github.com/rajithacharith/go-template/internal/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	err := config.LoadConfigs()
	if err != nil {
		logrus.Fatalf("Failed to load configs: %v", err)
	}

	configs := config.GetConfig()
	setLogLevel(configs.Server.LogLevel)

	app, err := app.MakeAPIServer()
	if err != nil {
		logrus.Fatalf("Failed to create app: %v", err)
	}

	go func() {
		logrus.WithField("port", configs.Server.Port).Info("Starting Hello World Service...")
		err := app.Listen(fmt.Sprintf(":%d", configs.Server.Port))
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

func setLogLevel(level string) {
	switch level {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}
