package main

import (
	"darklab_blog/blog/articles"
	"darklab_blog/blog/settings"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_cp"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_os"
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
		os.RemoveAll("build")
		os.MkdirAll("build", os.ModePerm)

		components := []Component{
			articles.NewHome(),
		}

		for _, comp := range components {
			comp.Write()
		}

		folders := utils_os.GetRecursiveDirs(settings.ProjectFolder)
		for _, folder := range folders {
			if utils_filepath.Base(folder) == "static" {
				utils_cp.Dir(folder.ToString(), utils_filepath.Join(settings.ProjectFolder, "build").ToString())
			}
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
}
