package main

import (
	"github.com/ds0nt/shed/pkg/api"
	"github.com/ds0nt/shed/pkg/log"
)

func main() {
	log.InitLogger()
	log.Info("Starting Shed")
	svc := api.NewService()
	svc.StartServer()
}
