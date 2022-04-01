package main

import (
	"app_es_service_go/bootstrap"
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library/closable"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	if err := bootstrap.Init(ctx); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-c
	cancel()
	time.Sleep(5 * time.Second)
	closable.Done()
}
