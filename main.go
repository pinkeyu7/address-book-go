package main

import (
	"address-book-go/config"
	"address-book-go/pkg/logr"
	"address-book-go/pkg/valider"
	"address-book-go/route"
	"flag"
	log "github.com/sirupsen/logrus"
)

var port string

func main() {
	flag.StringVar(&port, "port", "8080", "Initial port number")
	flag.Parse()

	config.InitEnv()

	// init logger
	logr.InitLogger()

	// validation init
	valider.Init()

	r := route.Init()

	// start server
	err := r.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
