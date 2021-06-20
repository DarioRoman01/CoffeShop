package main

import (
	protos "currecy/protos/currency"
	"currecy/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrencyServer(log)
	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	listener, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(listener)
}
