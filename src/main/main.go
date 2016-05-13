package main

import (
    "config"
	"log"
	"net/http"
	"apiserver"
)

func main() {
	config.LoadConfig()

	router := apiserver.NewRouter(routes)

	log.Fatal(http.ListenAndServe(config.Conf.Default.Bind, router))
}