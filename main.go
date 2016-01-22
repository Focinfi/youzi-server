package main

import (
	"github.com/Focinfi/apifaker"
	"log"
	"net/http"
	"os"
)

var goPath = os.Getenv("GOPATH")

var apisDir = goPath + "/src/github.com/Focinfi/youzi-server/public/fake_apis"

func main() {

	var addr string

	if len(os.Args) == 1 {
		addr = ":3000"
	} else if len(os.Args) == 2 {
		ip := os.Args[1]
		if ip != "" {
			addr = ip
		}
	} else {
		panic("Setting will be has only on param to points to ip.")
	}

	faker, err := apifaker.NewWithApiDir(apisDir)
	if err != nil {
		log.Fatal(err)
	} else {
		http.ListenAndServe(addr, faker)
	}
}
