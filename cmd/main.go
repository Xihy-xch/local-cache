package main

import (
	"cache/inject"
	"log"
)

func main() {
	app, cleanup, err := inject.InitApp(1024)
	if err != nil {
		log.Fatal("err: ", err)
	}

	if cleanup != nil {
		defer cleanup()
	}

	app.Start()
}
