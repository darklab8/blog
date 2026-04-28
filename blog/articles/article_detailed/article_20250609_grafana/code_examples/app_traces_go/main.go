package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/darklab8/go-typelog/otlp"
	"github.com/darklab8/go-utils/typelog"
	"go.opentelemetry.io/otel"
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

var (
	logger *typelog.Logger = typelog.NewLogger("go-demo-app")
	Tracer                 = otel.Tracer("go-demo-app")
)

func NestedAction(ctx_span context.Context) {
	ctx_span, span := Tracer.Start(ctx_span, "nested action")
	defer span.End()
}

func doRun() {
	time_start := time.Now()
	fmt.Println("started run", time_start)
	ctx_span, span := Tracer.Start(context.Background(), "web request")
	defer span.End()

	time.Sleep(3 * time.Second)

	web_endpoint := WebEndpoints[rand.IntN(len(WebEndpoints))]
	duration := rand.Float64() * web_endpoint.max_duration
	pattern := web_endpoint.pattern
	logger.InfoCtx(ctx_span, "web request",
		typelog.String("url_pattern", pattern),
		typelog.Float64("duration", duration),
		typelog.String("url_path", web_endpoint.url()),
	)
	NestedAction(ctx_span)
	fmt.Println("fninished run", time.Now(), time.Since(time_start))
	time.Sleep(3 * time.Second)
}

func main() {
	fmt.Println("starting app-traces")
	ctx := context.Background()
	otelShutdown, err := otlp.SetupOTelSDK(ctx) // Set up OpenTelemetry.
	if err != nil {
		fmt.Println("error to initialize tracing, err=", err.Error())
	}
	defer func() { // Handle shutdown properly so nothing leaks.
		err = errors.Join(err, otelShutdown(context.Background()))
	}()
	fmt.Println("configured tracing")
	for {
		doRun()
		time.Sleep(30 * time.Second)
	}
}
