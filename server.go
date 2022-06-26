package main

import "github.com/watshim-b/cln-at-go/controller"

func main() {
	r := controller.NewRouter()
	r.Run()
}
