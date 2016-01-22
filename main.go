package main

import (
	"github.com/Focinfi/apifaker"
	"log"
	"net/http"
	"os"
)

var pwd, _ = os.Getwd()

var developmentDir = pwd + "/public/fake_apis"

const productionDir = "/home/youzi-server/public/fake_apis"

func main() {

	var dir, addr string

	if len(os.Args) == 1 {
		dir = developmentDir
		addr = ":3000"
	} else if len(os.Args) == 3 {
		mode := os.Args[1]
		if mode == "production" {
			dir = productionDir
		}

		ip := os.Args[2]
		if ip != "" {
			addr = ip
		}
	} else {
		panic("Setting will be has a develop mode and a ip.")
	}
	log.Println(os.Args, dir, addr)

	faker, err := apifaker.NewWithApiDir(dir)
	if err != nil {
		log.Fatal(err)
	} else {
		http.ListenAndServe(addr, faker)
	}
}
