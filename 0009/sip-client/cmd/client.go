package main

import (
	"context"
	"flag"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"sip-client/pkg/instruction"
	_ "sip-client/pkg/util/log"
	_ "sip-client/pkg/util/timer"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	log.Info().Msgf("sip cli command started at: %s", time.Now().String())

	ctx, cancelFunc := context.WithCancel(context.Background())
	setupSignalHandler(cancelFunc)

	executor := instruction.NewExecutor(ctx)
	reader := instruction.NewReader(ctx)

	go executor.Run()
	go reader.Run()
	<-ctx.Done()
}

func setupSignalHandler(cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 2)
	shutdownSignals := []os.Signal{os.Interrupt, syscall.SIGTERM}
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		cancelFunc()
		<-c
		os.Exit(1)
	}()
}
