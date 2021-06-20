package main

import (
	"currecy/protos/currency"
	"currecy/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrencyServer(log)

	currency.RegisterCurrencyServer(gs, cs)
}
