package app

import (
	"fmt"
	server "github.com/avalonprod/gasstrem--back-end/internal/server"
	"log"
	"net/http"
)

const PORT string = "8000"

func HandlerApp() {
	server := new(server.Server)
	http.HandleFunc("/", homePage)
	if err := server.Run(PORT); err != nil {
		log.Fatalf("error accured whiule running http server: %v", err)
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server has been started...")
}
