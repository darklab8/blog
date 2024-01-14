package main

import (
	"darklab_blog/blog"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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

	build := func() {
		os.MkdirAll("build", os.ModePerm)

		components := []Component{
			blog.NewHome(),
		}

		for _, comp := range components {
			comp.Write()
		}
	}
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
		build()
	case ActionWeb:
		build()
		web()
	default:
		panic("action is not chosen")
	}

	// component := blog.Hello("John", 123)
	// component.Render(context.Background(), os.Stdout)

}
