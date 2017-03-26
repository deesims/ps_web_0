package main

import (
	"net/http"

	"ps_web_0/controller"
	"fmt"
)

//go:generate sqlboiler postgres --no-tests --wipe

const hostAddress = "localhost:7388"

func main() {
	controller.Init()

	fmt.Println("Starting server on " + hostAddress)
	http.ListenAndServe(hostAddress, nil)
}
