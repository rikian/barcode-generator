package main

import (
	"barcode/app/gin"
	"barcode/config"
	"barcode/db"
	"barcode/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config.LoadEnvFile()
	logger := config.BuildLogger()

	// create connection to sqlitedb
	dbConn := db.ConnectionToSql3()

	// initial repository
	repo := db.NewBarcodeRepository(dbConn, logger)

	// initial service
	service := service.NewServiceBarcode(logger, repo)

	// initial router
	router := gin.Router(logger, service)

	port := ":9091"

	// initial server
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		log.Print("Server running at " + port)
		logger.Info("Server running at " + port)

		// service connections
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutdown Server ...")
	logger.Info("Shutdown server ...")

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stop accepting new connections
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Print("timeout of 5 seconds.")
	}

	log.Print("Server exiting")
	logger.Info("Server exiting")
}
