package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type config struct {
	ENV string
}

func main() {
	log, err := initLogger() //start logger
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorf("error")
		os.Exit(1)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(1000)
	fmt.Println("Starting:", n)
	time.Sleep(5 * time.Second)
	fmt.Println("Ending:", n)
}

func run(log *zap.SugaredLogger) error {
	log.Infow("starting")
	var cfg config

	err := envconfig.Process("asd", &cfg)
	if err != nil {
		log.Errorw("error")
	}
	fmt.Println(cfg.ENV)
	api := http.Server{
		Addr:    ":5555",
		Handler: http.HandlerFunc(hello),
	}
	serverErrors := make(chan error, 1)
	go func() {
		log.Infow("main: Listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-serverErrors:
		log.Fatal(err)
	case <-shutdown:
		log.Infow("main: Start shutdown", "port", api.Addr)

		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err := api.Shutdown(ctx)
		if err != nil {
			log.Infow("main: Graceful shutdown did not complete in %d: %v", timeout, err)
			err = api.Close()
		}
		if err != nil {
			log.Fatalf("main: could not stop server gracefully: %v", err)
		}
	}
	return nil
}

func initLogger() (*zap.SugaredLogger, error) {
	logger, _ := zap.NewProduction()
	return logger.Sugar(), nil
}
