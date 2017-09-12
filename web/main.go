package main

import (
	"github.com/wizelineacademy/GoWorkshop/web/pkg/web"
	"log"
)

func main() {
	log.Print("[main] service started")
	web.ListenAndServe()
}
