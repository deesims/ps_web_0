package main

import (
	"net/http"

	"fmt"
	"github.com/deesims/ps_web_0/controller"
)

//go:generate sqlboiler postgres --no-tests --wipe

const hostAddress = "localhost:7388"

func main() {

	controller.Init()

	fmt.Println("Starting server on " + hostAddress)
	http.ListenAndServe(hostAddress, nil)
}
