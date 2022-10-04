package main

import (
	"context"
	"golang-simple-api-v2/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "golang-simple-api-v2", log.LstdFlags)

	// creates handlers
	productHandler := handlers.NewProducts(logger)

	// create a new serve mux and register the handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	// create a new server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel // Waiting for notification

	logger.Println("Recieved terminate, graceful shutdown", sig)

	deadline := time.Now().Add(30 * time.Second)
	timeoutContext, _ := context.WithDeadline(context.Background(), deadline)
	// Evita que al apagar el servidor no interrumpamos alguna acción que todavía no haya terminado
	// Y eso genere algún error. (Ejemplo: Transacción)
	server.Shutdown(timeoutContext)
}
