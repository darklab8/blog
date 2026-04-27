# Intro

in [a previous article]({{.SiteRoot}}article_grafana_loki.html) we configured Grafana, Loki and Alloy for logs gathering.

in this article we are going through configuring Tracing with Tempo, and we assume that previous configurations of a system were already done from the Part 1 of the article series (like Grafana web gui configuration and Caddy web server at least)

Tracing answers a question, if we have some long executed processes, how can we easily see which time each step takes? The same answer can be done with Profiling, but Profiling is not controllable and hard to navigate. The tracing is the best for Backend applications with many different network requests (to databases or http requests to apis and other your own services). Tracing has the next properties:
- Tracing is highly programmatic flexible in its configurations and serves as a glue between Logging, Metrics and Profiling.
- You can find Logs from Traces, Traces from Logs, Metrics from Traces, Traces from Metrics, Profiles from Traces. Traces link all monitoring for gliding between different perspectives to observe desired data
- Traces are less noisy than profiles and just have magnitudes better searching mechanism than Profiling to find the necessary data. The best working for debugging networking requests in backend applications
- As some disadvantage Traces has autoinstrumented enabling for automatic coverage only in specific languages for specific libraries
    - like in Python we have [instrumentation](https://opentelemetry-python-contrib.readthedocs.io/en/latest/) libs to turn on tracing automatically. Due to python otlp implementation being able to pass context in background by default, implementing high coverage is quite simple in python
    - Java should be having similar to Python easy propagation mechanics
    - But not in Golang. In golang we have to manually add Middlewares/handlers to libraries, and manually pass Context around nested functions in order to make that related Spans get connected.

{{ note `We have as some weak substistution for instrumentation in Go [epbf based tool](https://github.com/open-telemetry/opentelemetry-go-instrumentation), but it is highly limited, your logs, metrics will not have connections to traces, and it works only for specific sub set of libraries which u can't easily change. We will not be covering this tool usage in this series of article since it is not looking like good method to go by default.`}}

# Raising Tempo

{{ important `we provide docker-compose way of configuration as demo example because more devs are highly likely familiar and comfortable with docker-compose than with terraform. We utilize terraform for configuration of it and recommend it to use instead of docker-compose if u can. Book "Terraform up and running" is excellent place to start with it.` }}

<div class="code-container">
  <div class="code-tabs">
    <button class="tab-button active" data-lang="yaml">docker-compose.yaml</button>
    <button class="tab-button" data-lang="c">main.tf</button>
  </div>
  <pre class="code-block language-yaml"><code>
{{.TracingCompose}}
  </code></pre>
  <pre class="code-block language-c hidden"><code>
{{.MainTerraform}}
  </code></pre>
</div>

Participating configs:

{{ spoiler_start "Dockerfile.tempo" }}
{{ code_block_start "sh" }}
{{.TempoDockerfile}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "tempo.yaml" }}
{{ code_block_start "yaml" }}
{{.TempoConfig}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "Dockerfile.alloy.traces" }}
{{ code_block_start "sh" }}
{{.AlloyTracesDockerfile}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "cfg.traces.alloy" }}
{{ code_block_start "yaml" }}
{{.AlloyTracesConfig}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

Proceed to apply deployment for raising the tracing stack part
(or use Opentofu(Terraform) to raise all stuff together as modules from ./main.tf)
```sh
git clone --recurse-submodules https://github.com/darklab8/blog
cd blog/articles/article_detailed/article_20250609_grafana/code_examples

# if docker-compose way
export DOCKER_HOST=ssh://root@homelab
docker ps
docker compose -f docker-compose.tracing.yaml build
docker compose -f docker-compose.tracing.yaml up -d tempo
docker compose -f docker-compose.tracing.yaml up -d alloy-traces

# if opentofu way
tofu init
tofu apply

# after deploy, u need to grant tempo proper rights to be persistent and possible to init
chmod -R a+rw /var/lib/docker/volumes/tempo_data
```

{{ note `We presume "grafana" and "caddy" were raised in the previous part of the article about Loki. If that did not happen, start with it first there https://darklab8.github.io/blog/article_grafana_loki.html . raising grafana and caddy is fairly simple as "docker compose up -d grafana ; docker compose up -d caddy ; chmod -R a+rw /var/lib/docker/volumes/grafana_data. Code for their raising and configs in the same folder as docker-compose.tracing.yaml` }}
# Demo application to test it.

```sh
export DOCKER_HOST=ssh://root@homelab
docker compose -f docker-compose.app-traces.yaml build
docker compose -f docker-compose.app-traces.yaml run -it app-traces-go
```

with the next code is deployed
```go
package app_traces_go

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

  // example how to initialize Tracing itself is copy pasted from https://opentelemetry.io/docs/languages/go/getting-started/
  // and into https://github.com/darklab8/go-utils/blob/master/otlp/setup.go
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
```

and wee see in its logging its is working
```
> starting app-traces
> configured trading
> started run 2026-04-27 01:33:20.263967654 +0000 UTC m=+0.004603667
> fninished run 2026-04-27 01:33:23.264569558 +0000 UTC m=+3.005205571 3.000601995s
```

If everything is all right and no errors appears at any level in the chain of
- App works fine
- Grafana alloy works fine and has no errors regarding sending traces
- Tempo works fine and has no errors related to issues like unable to init backend because not having sufficient rights to initialize its data folder (to fix which u need to run `chmod -R a+rw /var/lib/docker/volumes/tempo_data/`)
- Grafana works fine and initialized itself with provisioning data resources
- Grafana plugin for tracing drilldown works fine as well

You will see traces visible in your tracing drilldown interface then!

![]({{.StaticRoot}}grafana_tempo/tempo2.png)
![]({{.StaticRoot}}grafana_tempo/tempo3.png)

In a real world tracing is the most useful for backend applications and the best to turn it on by default for all the network interacting libraries through writing some kind of middleware.
Then it will be able to answer you that issues you have at specific SQL request, or elastic search query, or specific http request. And since it is distributed tracing, the trace will shown how workload works within the called service too!

{{ note `In pet projects Tracing usability is honestly very limited, since it is very doubtful for pet project to have any kind of network interaction long enough requiring tracing debugging. Just because your database will rarely reach the level requiring to debug it. You will benefit in pet projects more from Logging and Metrics monitoring system. At any backend real work though, tracing is the most useful system to have, i would dare to say potentially even more useful than any other type of monitoring.` }}