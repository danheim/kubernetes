package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/takama/router"
)

var log = logrus.New()

func main() {
	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		log.Fatal("Required parameter service port is not set")
	}

	r := router.New()
	r.Logger = logger
	r.GET("/", home)
	r.Listen(":" + port)
}
