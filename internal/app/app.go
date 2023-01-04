package app

import (
	server "github.com/avalonprod/gasstrem--back-end/internal/server"
	"log"
)

const PORT string = "8000"

func HandlerApp() {
	server := new(server.Server)
	if err := server.Run(PORT); err != nil {
		log.Fatalf("error accured whiule running http server: %v", err)
	}
}
