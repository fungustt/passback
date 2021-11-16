package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"passback/cmd/pass/generator"
	"passback/cmd/pass/http"
	"passback/cmd/pass/http/handlers"
)

func main() {
	ctx := initGraceful()

	healthHandler := handlers.NewHealthHandler()

	gen := generator.NewGenerator(ctx, 10)
	passwordHandler := handlers.NewPasswordHandler(gen)

	server := http.NewServer(":10000", []http.Handler{healthHandler, passwordHandler})
	server.Start(ctx)
}

func initGraceful() context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-c
		log.Printf("system call: %+v", osCall)
		cancel()
	}()

	return ctx
}
