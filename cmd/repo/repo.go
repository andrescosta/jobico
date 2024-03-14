package main

import (
	"log"

	"github.com/andrescosta/goico/pkg/context"
	"github.com/andrescosta/jobico/cmd/repo/service"
	_ "github.com/tprasadtp/go-autotune"
)

func main() {
	ctx, cancel := context.ForEndSignals()
	defer cancel()
	svc, err := service.New(ctx)
	if err != nil {
		log.Panicf("error creating queue service: %s", err)
	}
	defer svc.Dispose()
	if err := svc.Start(); err != nil {
		log.Panicf("error starting ctl service: %s", err)
	}
}
