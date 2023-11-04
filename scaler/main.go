package main

import (
	"flag"
	"log"
	"scaler/common/logger"
	"scaler/common/public"
	"scaler/internal/server"
)

func main() {
	cfg := flag.String("c", "config.yaml", "config file path.")
	flag.Parse()

	public.SetHandler(*cfg)
	if err := logger.Setup(public.GetHandler().Config.Server.Loglevel); err != nil {
		log.Panicf("logger setup failed: %s", err.Error())
	}
	if err := server.New().Start(); err != nil {
		log.Panicf("start server failed: %s", err.Error())
	}
}
