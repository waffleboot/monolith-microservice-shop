package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Context() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
			return
		}
	}()
	return ctx
}
