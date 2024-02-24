package app

import (
	"flag"
	"github.com/maxzhovtyj/card-validator/internal/config"
	"github.com/maxzhovtyj/card-validator/internal/handler/handler"
	"github.com/maxzhovtyj/card-validator/internal/service"
	"github.com/maxzhovtyj/card-validator/pkg/card"
	"github.com/maxzhovtyj/card-validator/pkg/log/applogger"
	"google.golang.org/grpc"
	"net"
)

var configPath = flag.String("configPath", "", "path to config file")

func Run() {
	logger := applogger.New()

	logger.Infof("init application config")
	flag.Parse()

	if *configPath == "" {
		logger.Fatalf("empty config path")
	}

	cfg, err := config.New(*configPath)
	if err != nil {
		logger.Fatal(err)
	}
	cfg.LogAll()

	logger.Infof("init tcp listener")
	tcp, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infof("init card validator")
	validator := card.NewValidator()

	logger.Infof("init services")
	s := service.New(validator)

	logger.Infof("init grpc server")
	srv := grpc.NewServer()
	handler.NewServer(srv, s, logger)

	logger.Infof("start grcp server on address '%s'", cfg.Addr)
	if err = srv.Serve(tcp); err != nil {
		logger.Fatal(err)
	}
}
