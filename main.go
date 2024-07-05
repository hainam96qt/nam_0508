package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"nam_0508/internal/config"
	"nam_0508/internal/endpoint/purchase"
	wager2 "nam_0508/internal/endpoint/wager"
	purchase3 "nam_0508/internal/repo/purchase"
	wager3 "nam_0508/internal/repo/wager"
	purchase2 "nam_0508/internal/service/purchase"
	"nam_0508/internal/service/wager"
	"nam_0508/pkg/db/mysql_db"
)

func main() {
	startedAt := time.Now()
	defer func() {
		log.Printf("application stopped after %s\n", time.Since(startedAt))
	}()

	conf, err := configs.NewConfig()
	if err != nil {
		log.Print(err)
	}

	globalCtx, glbCtxCancel := context.WithCancel(context.Background())

	httpSrv, err := initHTTPServer(globalCtx, conf)
	if err != nil {
		log.Panicf("failed to init http server %s \n", err)
	}

	go func() {
		log.Printf("starting HTTP server at: %s\n", conf.Server.Address)
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("failed to start HTTP server: %s \n", err)
		}
	}()

	// Keep the application running until signals trapped
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("%s signal trapped. Stopping application", <-sigChan)

	glbCtxCancel()
	// First terminate the HTTP gateway
	shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), conf.Server.ShutdownTimeout)
	defer shutdownCtxCancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		log.Printf("failed to gracefully shutdown the HTTP gateway server: %s\n", err)
	} else {
		log.Println("HTTP gateway server stopped gracefully")
	}
}

func initHTTPServer(ctx context.Context, conf *configs.Config) (httpServer *http.Server, err error) {
	r := chi.NewRouter()

	// create endpoint here
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err = w.Write([]byte("welcome"))
	})

	dbConn, err := mysql_db.ConnectDatabase(conf.Mysqldb)
	if err != nil {
		log.Panicf("failed to connect database:: %s \n", err)
		return
	}

	// repository
	purchaseRepo := purchase3.NewMysqlRepository(dbConn)
	wagerRepo := wager3.NewMysqlRepository(dbConn)

	// service
	wagerSvc := wager.NewWagerService(wagerRepo)
	purchaseSvc := purchase2.NewPurchaseService(dbConn, purchaseRepo, wagerRepo)

	// handler
	wager2.InitWagerHandler(r, wagerSvc)
	purchase.InitPurchaseHandler(r, purchaseSvc)

	return &http.Server{
		Addr:         conf.Server.Address,
		ReadTimeout:  conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
		Handler:      r,
	}, nil
}
