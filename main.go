package main

import (
	"flag"
	"github.com/Focinfi/apifaker"
	"log"
	"net/http"
	"os"
)

var goPath = os.Getenv("GOPATH")

var apisDir = goPath + "/src/github.com/Focinfi/youzi-server/public/fake_apis"

func main() {
	addr := flag.String("b", ":3000", "bind a ip with port")
	flag.Parse()

	faker, err := apifaker.NewWithApiDir(apisDir)
	if err != nil {
		log.Fatal(err)
	} else {
		http.ListenAndServe(*addr, faker)
	}
}
