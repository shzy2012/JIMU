package main

import (
	"context"
	"flag"
	"jimu/src/db"
	"jimu/src/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/shzy2012/common/log"
)

func init() {
	db.Init()
}
func main() {

	log.Println("[server]=>starting.")
	var port int
	flag.IntVar(&port, "p", 80, "http 端口号  -p=80")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go service.Nicedog()
	go service.Serve(ctx, port)

	go func() {
		sig := <-sigs
		log.Info(sig)
		log.Println("Shuting down server...")
		cancel()
		done <- true
	}()

	<-done
	log.Info("[server]=>stop service.")
}
