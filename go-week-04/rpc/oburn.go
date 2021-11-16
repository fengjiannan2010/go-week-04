package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"oburn/rpc/cronjob/logic"
	"oburn/rpc/internal/config"
	"oburn/rpc/internal/server"
	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"
)

var configFile = flag.String("f", "etc/oburn.yaml", "the config file")

func main() {
	flag.Parse()
	rootContext := context.Background()
	cctx, cancelFunc := context.WithCancel(rootContext)
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	db, _ := ctx.OrmDb.DB()
	defer db.Close()

	logic.InitLogic(cctx, ctx)

	srv := server.NewOburnServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		oburn.RegisterOburnServer(grpcServer, srv)
		reflection.Register(grpcServer)
	})
	defer s.Stop()
	go logic.Task.Start()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
	cancelFunc()
}
