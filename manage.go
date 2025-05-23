package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/darklab8/blog/blog"
	"github.com/darklab8/blog/blog/settings"
	"github.com/darklab8/blog/blog/settings/logus"
	"github.com/darklab8/go-utils/utils/cantil"
)

func main() {
	web := func() {
		fs := http.FileServer(http.Dir("./build"))
		http.Handle("/", fs)

		port := 8000
		logus.Log.Info("Listening on " + strconv.Itoa(port))
		err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
		logus.Log.CheckPanic(err, "unable to serve serve")
	}

	parser := cantil.NewConsoleParser(
		[]cantil.Action{
			{
				Nickname:    "build",
				Description: "build to static assets: html, css, js files",
				Func: func(info cantil.ActionInfo) error {
					blog.Builder.BuildAll()
					return nil
				},
			},
			{
				Nickname:    "web",
				Description: "for local development",
				Func: func(info cantil.ActionInfo) error {
					blog.Builder.BuildAll()
					web()
					return nil
				},
			},
		},
		cantil.ParserOpts{
			Enverants: settings.Enverants,
		},
	)
	parser.Run(os.Args[1:])
}
