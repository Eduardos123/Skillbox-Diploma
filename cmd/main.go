package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/DwarfWizzard/skillbox-diplom-go/internal/service"
	"github.com/gin-gonic/gin"
)

const pathToData = "../data"
const uri = "http://127.0.0.1:8383"
const host = "127.0.0.1:8585"

func main() {
	ctx := context.Background()

	svc := service.New(pathToData, uri)
	router := gin.New()
	router.GET("/", svc.HandleConnection)

	server := http.Server{Addr: host}
	server.Handler = router

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-quit

	log.Println("shutdown server")
	if err := server.Shutdown(ctx); err != nil {
		log.Print(err)
	}
}
