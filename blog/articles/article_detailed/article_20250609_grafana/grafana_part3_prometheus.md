# Intro

Monitoring with metrics is the most useful because they tell you roughly when and where is the issue. 
- For example we can see CPU/RAM/disk usage and other general metrics about application running, including when there were potentially reported OOM kills, with history of data in period of for example last month
- Or we could be seeing detalization, regarding which our endpoints where slow, which percentage of them were slow? At which endpoint web server spent the most sum of time (after we summed time of all processed responses)?
- Or we could be seeing where our message queue handling workers are stuck, where they error and etc.

Metrics show detalization When, and What. historic data that gives us understanding where is the source of issue or timeframe where to dig for data (by logs and traces) further.

You no longer need real time watching current values of application usage, you are able monitor historic data and finding clues among the noticed patterns!

The most important quality of Metrics is that they are HIGHLY PERFORMANT to query and to store. they are taken one time per 1 minute for example, per 1 unit set of labels (application_name, endpoint_name, status_code as example of labels). So when Logging monitoring we can be struggling to have weeks data retention, and tracing is barely affordable to see days of data, we are still able to see **Metrics data for many months even in the most high loaded infrastructure.**

## The most common issue. High cardinality.

The biggest and most often mistake people make is having a lot of unique labeled values in some metric. that causes rapid RAM usage explosion and storage as well.
You can debug which metric consumes a lot of space by using query
```
topk(20, count by (__name__, job)({__name__=~".+"}))
```

optionally if you know specific label in a metric has a lot of values but you have trouble to identify which ones, you can be outputing metric label values grouped by first N characters.
Here is example we outputed span_name in traces_spanmetrics_latency_count, grouped by first 5 characters.
```
sum by (span_prefix) (
  label_replace(
    traces_spanmetrics_latency_count,
    "span_prefix",
    "$1",
    "span_name",
    "(.{5}).*"
  )
)
```

##  Raising Prometheus

{{ important `we provide docker-compose way of configuration as demo example because more devs are highly likely familiar and comfortable with docker-compose than with terraform. We utilize terraform for configuration of it and recommend it to use instead of docker-compose if u can. Book "Terraform up and running" is excellent place to start with it.` }}

<div class="code-container">
  <div class="code-tabs">
    <button class="tab-button active" data-lang="yaml">docker-compose.yaml</button>
    <button class="tab-button" data-lang="c">main.tf</button>
  </div>
  <pre class="code-block language-yaml"><code>
{{.MetricsCompose}}
  </code></pre>
  <pre class="code-block language-c hidden"><code>
{{.MainTerraform}}
  </code></pre>
</div>

Participating configs:

{{ spoiler_start "prometheus.yaml" }}
{{ code_block_start "yaml" }}
{{.MetricsConfig}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "cfg.metrics.alloy" }}
{{ code_block_start "yaml" }}
{{.MetricsAloyConfig}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "Dockerfile.prometheus" }}
{{ code_block_start "sh" }}
{{.MetricsDockerfile}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "Dockerfile.metrics.traces" }}
{{ code_block_start "sh" }}
{{.MetricsAlloyDockerfile}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

Proceed to apply deployment for raising the metrics stack part
(or use Opentofu(Terraform) to raise all stuff together as modules from ./main.tf)
```sh
git clone --recurse-submodules https://github.com/darklab8/blog
cd blog/articles/article_detailed/article_20250609_grafana/code_examples

export DOCKER_HOST=ssh://root@demo
docker ps

# ONLY if you did not do things from previous article part about Loki or Tempo follow docker-compose path:
docker compose up -d caddy # we need it for reverse proxy and automated TLS certs
docker compose up -d grafana # visualizer where we query traces. Already yaml of provisioned datasources configured

# Continue with Prometheus article content:
# if docker-compose way:
docker compose -f docker-compose.prometheus.yaml build
docker compose -f docker-compose.prometheus.yaml up -d prometheus
docker compose -f docker-compose.prometheus.yaml up -d alloy-metrics

# if opentofu way
tofu init
tofu apply

# after deploy, u need just in case to grant prometheus proper rights to be persistent and possible to init
chmod -R a+rw /var/lib/docker/volumes/prometheus_data
chmod -R a+rw /var/lib/docker/volumes/grafana_data # just in case grant grafana rights too if not granted
```

If everything was configured correctly, you will be able to open Metrics Drilldown page and see incoming metrics already. This article brings alloy configuration with prewritten docker monitoring, as it is the most comfortable minimalistic approach for deployment in a homelab.

![]({{.StaticRoot}}grafana_prometheus/metrics_drilldown.png)

{{ note `If you wish monitoring by metrics something else besides docker and applications in docker, for example postgres, elasticsearch, aws cloudwatch and etc. Check other grafana alloy components ( https://grafana.com/docs/alloy/latest/reference/components/prometheus/ ) for other provided prometheus integrations` }}

## Dashboards

We raised in previous section prometheus and universal metrics metrics scrapper "alloy", which is already configured to scrap unix, docker and app metrics.
- [Cadvisor alloy component](https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.exporter.cadvisor/) for docker metrics
- [Unix alloy component](https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.exporter.unix/) for node exporter metrics accordingly

![]({{.StaticRoot}}grafana_prometheus/import_dashboard.png)

Import dashboards for Docker and Unix
- [node exporter dashboard json](https://raw.githubusercontent.com/darklab8/infra/refs/heads/master/tf/grafana/grafana_stack/dashboards_unix/NodeExporterBasedOn11074.json)
  - [import from original source](https://grafana.com/grafana/dashboards/15172-node-exporter-for-prometheus-dashboard-based-on-11074/)
- [cadvisor, docker monitoring dashboard json](https://raw.githubusercontent.com/darklab8/infra/refs/heads/master/tf/grafana/grafana_stack/dashboard_docker/cadvisor_exporter_14282_rev1.json)

If you imported all right and your Grafana image version is 11.6 as supported by those dashboards and written in docker compose, you will see metrics about your containers and linux server accordingly

Cadvisor dashboard (about docker containers):
![]({{.StaticRoot}}grafana_prometheus/dashboard_cadvisor.png)
Node exporter dashboard (about linux server):
![]({{.StaticRoot}}grafana_prometheus/dashboard_node_exporter.png)

{{ critical `if you remain not seeing dashboards properly, take note which grafana version you use. We can be sure it works fine with 11.6 at least` }}

{{ note `grafana side of configurations author of article handles by terraform grafana provider instead of manual actions. https://github.com/darklab8/infra/tree/master/tf/grafana` }}

if you what other kind of grafana dashboards, you could browse all choices people release there https://grafana.com/grafana/dashboards/ , of you can make your own.

{{ critical `if imported dashboards do not show data also in some of its graphs, make sure you have in data source of prometheus turned on timeInterval to 60s, the time of alloy scraping interval. If u raised grafana web interface with this article things, you will have it automatically as it is written in data source provisioning config` }}

# Application dashboards

Now that we have main infra dashboards handled, we now need to try having some custom application metrics scraped and made into its own dashboard
- Supported languages for prometheus libraries can be found here https://prometheus.io/docs/instrumenting/clientlibs/
- Depending on a language, framework, infra element, there can be already existing integrations/exporters for it, which u could find here https://prometheus.io/docs/instrumenting/exporters/ , like there is even Python django prometheus integration and postgresql integrations
- As mentioned before grafana alloy scraping agent offers plenty of common exporters u could choose to use to get access to more metrics https://grafana.com/docs/alloy/latest/reference/components/prometheus/ . to use them properly u will highly likely need reading their real repository under the hood about necessary extra volumes/settings u need to pass to grafana alloy to make it working depending on what you use

Read throughly metric types existing there https://prometheus.io/docs/concepts/metric_types/ to understand how write your own prometheus metrics. Roughly we can say
- Counter is good for stuff like request counts, action counts, well any kind of counts most of the time.
- Gauge for If we need to know "Temperature value" of smth, how many workers we have currently or how many active users are currently if we have access to active sessions. Also gauge is usable for summing up anything.
- Histogram when we need to capture performance of request duration, or any other kind of duration across different "route patterns".

I integrated my pet project with Go lib of Prometheus https://github.com/prometheus/client_golang
added metrics and registered them in explicit way for ability to add global labels
https://github.com/darklab8/fl-darkstat/blob/master/darkcore/metrics/metrics.go

Based on that I have for project darkstat detailed performance evaluating dashboard https://grafana.dd84ai.com/d/belbdnu2uqe4gd/app-darkstat?var-interval=2m&orgId=1&from=now-3h&to=now&timezone=browser&var-environment=production

- Since it is a web app, i made sure to capture regular stuff of how success/failure rates and duration of responses my web server makes
- And i had plenty problems with uptime in the past, so i made sure i have uptime dashboard present working through a regular Counter

![]({{.StaticRoot}}grafana_prometheus/app1.png)

- I took extra notice in worst case scenaries, when worst 50%,25%,10% of responses happen? (P50,P75,P90 metrics)
- I took notice which page are having large body size and create large network trafic for me. Due to nature of the app it was important for me to look for this metric.
- I took notice which pattern path take the most time to load to mark then as potential targets for optimization

![]({{.StaticRoot}}grafana_prometheus/app2.png)

- Lastly i keep running for the app public API, so out of curiosity i watch which endpoints are actually in use, with which user agents to evaluate amount of users
- That information gives me information what is NOT in use and i could consider evaluating to be removed as not necessary.

![]({{.StaticRoot}}grafana_prometheus/app3.png)

The code of this application dashboard is [provided by link](https://github.com/darklab8/infra/blob/master/tf/grafana/grafana_stack/dashboards_apps/darkstat.json)

What we capture in application dashboards?

We capture with metrics what is most important for us to monitor:
- how responses are given by web server
- How interactions with databases work inside the app
- how requests to other applications behave
- how payments are processed.
- how message queue workers work
- how databases run and behave, handling pressure, ram, disk usage

The role of dashboard metrics is to tell us WHERE the issues are happening (but they do not have to tell exactly how they are happening, see traces/logs/profiles for more information). Plus metrics are also the most useful to be utilized in Alerts because of their performance efficiency :)

# Dashboards based on logs and traces

if we haven't mentioned it before, we can make dashboards even from logs, but they will not be query performant and their usability is limited to applications with low logging volume. It is way easier to handle prometheus metrics that are emitted fews per minutes than applications with thousands of log lines per minute.

Example of dashboard based on logs can be found [here](https://grafana.dd84ai.com/d/feoxz03mpa39ce/darkstat-uses-logs?orgId=1&from=now-1h&to=now&timezone=browser)  and its [code is here](https://github.com/darklab8/infra/blob/master/tf/grafana/grafana_stack/dashboards_apps/darkstat_from_logs.json)

Dashboard from logs:

![]({{.StaticRoot}}grafana_prometheus/app_from_logs.png)

As for dashboards based on traces?
We can make generic dashboard that works based on metrics generated from traces. It is [provided here](https://grafana.dd84ai.com/d/tempo-metricsgenerator-front/trace-metrics-by-metricsgenerator?var-interval=10m&var-p_latency=95&orgId=1&from=now-1h&to=now&timezone=browser&var-service=alloy&var-span_name=.%2A&var-span_kind=$__all&refresh=30s) and its [code is here](https://github.com/darklab8/infra/blob/master/tf/grafana/grafana_stack/tracemetrics/tracementrics.json)

![]({{.StaticRoot}}grafana_prometheus/app_from_traces.png)

You may found it interesting because u get auto generated dashboard just because your app is connected to tracing. It has strong limitations of needing low cardinality used tracing span names though, if some application breaks this rule, it needs to be excluded from metrics generator usage.

{{ critical `turn off metrics generator is you are not needing trace apm dashboard, that will save you some used RAM` }}

Additionally dashboard graphs can be even generated from [traceql metrics](https://grafana.com/docs/tempo/latest/metrics-from-traces/metrics-queries/), which can be useful in tricky tracing searches. They aren't again useful for average every day usage because performance demand is too high there. Only Metrics based dashboards are performance efficient to be navigated plentifully.

# Articles updates

All articles about monitoring configurations, including about prometheus metrics, are actively in use at least in homelab of the author through Terraform configuration. If you have any doubts, smth got outdated, less working and etc, see [terraform code there as source of truth](https://github.com/darklab8/infra/tree/master/tf/modules/docker_stack/monitoring). With some chance article content will be updated in [its repository](https://github.com/darklab8/blog) and redeployed to github pages with fresh fixes. I can be reached for communications in [issues to blog repo](https://github.com/darklab8/blog/issues).

# Production grade configuration tips.

While prometheus is fine for few hosts of homelab, or very small production:
- I can highly endorce deploying Mimir if you have serious production in horizontal scaled infrastructure of kubernetes
https://github.com/grafana/mimir/tree/main/operations/helm/charts/mimir-distributed
The main advantage of it... it is actually horizontally scalable and able to withstand higher workload just because it can distribute RAM workload between its scaled instances.
- in its turn [k8s monitoring](https://github.com/grafana/k8s-monitoring-helm) helm chart is perfect in scraping metrics from kubernetes. Grafana alloy is perfect for usage in kubernetes, or in dockerized deployments, or in AWS ECS, or even for deployment to linuxes that have all stuff running through systemd, it can work anywhere and scraping everything.
- To make Mimir working for full power, you could be wishing to learn how to configure [Mimir Rules](https://registry.terraform.io/providers/fgouteroux/mimir/latest/docs/resources/rules). That will unlock you having in Mimir detailedly working dashboards of any kind if u just import the same Prometheus Rules as provided for whatever dashboards.
- Some people choose Victoria Metrics as alternative. It may be good alternative, but author of the article did not test it on a workload comparable to Mimir to be sure which one choice works better. Since there are no serious complains to Mimir and it works pretty fine and scalable, there was no yet need to switch to Victoria Metrics for the author.
