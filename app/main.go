package main

import (
	"app/common/logger"
	"app/common/public"
	"app/internal/server"
	"flag"
	"log"
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
