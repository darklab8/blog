package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/darklab8/blog/blog"
	"github.com/darklab8/blog/blog/settings/logus"
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

		port := 8080
		logus.Log.Info("Listening on " + strconv.Itoa(port))
		err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
		logus.Log.CheckPanic(err, "unable to serve serve")
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
