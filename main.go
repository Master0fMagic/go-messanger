package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-messanger/config"
	"go-messanger/server/http"
	"go-messanger/server/http/handler"
	"go-messanger/service/postgres"
	"go-messanger/service/postgres/provider"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	setupGracefulShutdown(cancel)

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err) //todo change to fatalf
	}

	pg, err := postgres.New(cfg.PostgresConfig)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := pg.Close()
		if err != nil {
			panic(err)
		}
	}()

	accountProvider := provider.NewAccountProvider(pg)
	accountHandler := handler.NewAccountHandler(accountProvider)

	srv := http.NewServer(cfg.HttpConfig, *accountHandler)

	go func() {
		err := srv.Run()
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		<-ctx.Done()
		wg.Done()
	}()

	wg.Wait()
}

func setupGracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		log.Warnf("Got Interrupt signal: %v", sig.String())
		stop()
	}()
}
