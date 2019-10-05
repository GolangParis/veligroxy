package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/GolangParis/veligroxy/internal/routes"
	"github.com/GolangParis/veligroxy/internal/version"
)

func main() {
	logger := logrus.New().WithField("Version", version.Version).WithField("Commit", version.Commit)
	logger.Info("application is starting...")

	// reading ports
	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Running port is not specified")
	}
	diagPort := os.Getenv("DIAG_PORT")
	if diagPort == "" {
		logger.Fatal("Diagnostic port is not specified")
	}

	r := routes.BusinessRoutes()
	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}

	// diag routes
	diagRouter := routes.DiagnosticsRoutes()
	diag := http.Server{
		Addr:    net.JoinHostPort("", diagPort),
		Handler: diagRouter,
	}

	// error channel
	shutdown := make(chan error, 2)

	// async start
	go func() {
		logger.Info("Business server is preparing...")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			shutdown <- err
		}
	}()

	go func() {
		logger.Info("Diagnostics server is preparing...")
		err := diag.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			shutdown <- err
		}
	}()

	// graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		logger.Warnf("Received `%v`. Application stopped.", x)

	case err := <-shutdown:
		logger.Warnf("Received shudown message: `%v`", err)
	}

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err := server.Shutdown(timeout)
	if err != nil {
		logger.Error(err)
	}

	errDiag := diag.Shutdown(timeout)
	if err != nil {
		logger.Error(errDiag)
	}

}
