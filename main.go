package main

import (
	"address-book-go/config"
	"address-book-go/pkg/valider"
	"address-book-go/route"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var port string

func main() {
	flag.StringVar(&port, "port", "8080", "Initial port number")
	flag.Parse()

	config.InitEnv()
	valider.Init()

	fmt.Println(config.GetEnvironment())

	r := route.Init()

	// start server
	err := r.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
