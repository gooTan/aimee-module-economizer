package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/JBailes/aimee/server-go/bus"
	handler "github.com/JBailes/aimee/server-go/modules/economizer"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s DAEMON_MODULE_BUS_SOCKET\n", os.Args[0])
		os.Exit(2)
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	config := bus.ModuleProcessConfig{
		SocketPath: os.Args[1], ModuleName: "economizer",
		PrincipalClass: 1, PrincipalRef: 27,
		Stages: []bus.ModuleStage{
		{EventKind: 11009, StageID: 1},
		{EventKind: 11010, StageID: 2},
		{EventKind: 11011, StageID: 3},
		{EventKind: 11012, StageID: 4},
		{EventKind: 11013, StageID: 5},
		},
		Handler: handler.Handle,
	}
	if err := bus.RunModuleProcess(ctx, config); err != nil {
		fmt.Fprintf(os.Stderr, "aimee-module-economizer: %v\n", err)
		os.Exit(1)
	}
}
