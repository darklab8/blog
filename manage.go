package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/darklab8/blog/blog"
)

type Component interface {
	Write()
}
type Action string

const (
	ActionWeb   Action = "web"
	ActionBuild Action = "build"
)

func main() {
	var action string
	flag.StringVar(&action, "act", "undefined", "action to run")
	flag.Parse()

	fmt.Println("act:", action)

	web := func() {
		fs := http.FileServer(http.Dir("./build"))
		http.Handle("/", fs)

		log.Print("Listening on :8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	switch Action(action) {
	case ActionBuild:
		blog.Builder.BuildAll()
	case ActionWeb:
		blog.Builder.BuildAll()
		web()
	default:
		panic("action is not chosen")
	}
}
