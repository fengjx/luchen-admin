package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengjx/luchen"

	"github.com/fengjx/luchen-admin/service"
	"github.com/fengjx/luchen-admin/transport"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := luchen.Logger(ctx)
	logger.Info("app start")
	service.Init()
	transport.Start(ctx)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	<-quit
	logger.Info("app stop")
	cancel()
	transport.Stop(ctx)
}
