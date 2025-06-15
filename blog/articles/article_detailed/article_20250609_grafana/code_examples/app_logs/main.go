package main

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
	"time"
)

type WebEndpoint struct {
	pattern      string
	max_duration float64
	url          func() string
}

var WebEndpoints = []WebEndpoint{
	{
		pattern:      "/index.html",
		max_duration: 0.1,
		url:          func() string { return "/index.html" },
	},
	{
		pattern:      "/some_pattern1",
		max_duration: 1,
		url:          func() string { return "/some_pattern1" },
	},
	{
		pattern:      "/another_pattern",
		max_duration: 2,
		url:          func() string { return "/another_pattern" },
	},
	{
		pattern:      "/books/__book_id__",
		max_duration: 0.1,
		url:          func() string { return fmt.Sprintf("books/%d", rand.IntN(100)) },
	},
	{
		pattern:      "/books/__book_id__/page/__page_id__",
		max_duration: 0.2,
		url:          func() string { return fmt.Sprintf("books/%d/page/%d", rand.IntN(100), rand.IntN(1000)) },
	},
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

/*
Application imitating web application with some traffic
*/
func main() {
	for {
		web_endpoint := WebEndpoints[rand.IntN(len(WebEndpoints))]
		duration := rand.Float64() * web_endpoint.max_duration
		pattern := web_endpoint.pattern
		logger.Info("web request",
			slog.String("url_pattern", pattern),
			slog.Float64("duration", duration),
			slog.String("url_path", web_endpoint.url()),
		)
		time.Sleep(5 * time.Second)
	}
}
