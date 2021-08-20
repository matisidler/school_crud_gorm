package main

import (
	"log"
	"net/http"
	"school/handlers"
)

func main() {

	mux := http.NewServeMux()
	handlers.Route(mux)

	log.Println("initialized server in 8050 port")
	err := http.ListenAndServe(":8050", mux)

	if err != nil {
		log.Println("server error: ", err)
	}

}
