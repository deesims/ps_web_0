package main

import (
	"net/http"

	"github.com/deesims/ps_web_0/controller"
)

//go:generate sqlboiler postgres --wipe

func main() {
	controller.Init()
	http.ListenAndServe("localhost:7388", nil)
}
