package main

import (
	"address-book-go/api"
	"address-book-go/config"
	"address-book-go/pkg/logr"
	"address-book-go/pkg/valider"
	"address-book-go/route"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var port string

// @title Address book API
// @version 1.0
// @description Address book API
// @termsOfService https://github.com/pinkeyu7/address-book-go
// @license.name MIT
// @license.url
func main() {
	// init http port
	flag.StringVar(&port, "port", "8080", "Initial port number")
	flag.Parse()

	// init config
	config.InitEnv()

	// init logger
	logr.InitLogger()

	// init validation
	valider.Init()

	// init driver
	_ = api.InitXorm()
	_ = api.InitRedis()

	// init gin router
	r := route.Init()

	// start server
	err := r.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
