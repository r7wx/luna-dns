package main

import (
	"os"

	"github.com/r7wx/luna-dns/internal/config"
	"github.com/r7wx/luna-dns/internal/engine"
	"github.com/r7wx/luna-dns/internal/logger"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		logger.Fatal("no configuration file provided")
	}
	config, err := config.Load(args[0])
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Configuration file loaded: " + args[0])

	engine, err := engine.NewEngine(config)
	if err != nil {
		logger.Fatal(err)
	}
	if err := engine.Start(); err != nil {
		logger.Fatal(err)
	}
}
