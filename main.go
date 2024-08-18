package main

import (
	"log"

	"github.com/nanpipat/nanpipat-agnos-backend/cmd"
	"github.com/nanpipat/nanpipat-agnos-backend/internal/config"
)

func main() {
	cfg := config.New()
	switch cfg.Service {
	case "api":
		cmd.APIRun()
	default:
		log.Fatal("Service not supported")
	}
}
