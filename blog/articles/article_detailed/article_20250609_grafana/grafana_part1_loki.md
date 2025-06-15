# Intro

Grafana monitoring stack with Alloy, Loki, Prometheus, Tempo is distributed modern monitoring made in Go, intended for collecting monitoring information from backend applications running at one or more servers. It is also usable for mobile / desktop applications monitoring (depends on things like level of Opentelemetry support for your chosen language if we go with otlp at least. For example [for Java opentelemetry reports Stable level for both Traces, Metrics and Logs](https://opentelemetry.io/docs/languages/java/), and thus full support is available for Android java apps), and to some degree usable for web frontend client side monitoring too.

{{ note `The article is written with assumption we will be using Opentelemetry protocol as main one for tracing related activities at least (but also abusing for metrics and logs too if necessary), but do know that Grafana Alloy tracing/logging/metrics collecting agent supports plenty of alternative protocols and your language support in them could be with some chance better than in opentelemetry.` }}

In this series of articles we will overview configuring this monitoring stack in Docker based approach for your homelab and for companies with simplistic approach to infrastructure. The article is aimed to make the monitoring system more accessable for major amount of people (their homelabs and basic production setups) and for this reason we go with Docker approach instead of Kubernetes one. If you run serious production with high load, it is better running Grafana/Loki/Mimir(instead of Prometheus)/Tempo in a Kubernetes instead, since its ecosystem with helm charts already made easy to run it in horizontal scalable way and able to take much larger workload.
The article will dive into configuring the monitoring with Docker-Compose and Opentofu(Terraform). When in doubts regarding smth, check Terraform related code for source of truth.

It is worth to configure this monitoring distributed stack even you have only one backend application running at your servers (or even if u have only mobile app). Well configured monitoring will grant you ability to debug your application information significantly easier. Well configured logging monitoring backend wil give you ability to filter data by any key/value in the logging records.
It is possible even build graphical dashboards based on Logging Information alone for overview of important information!

{{ note `Grafana Loki became significantly more pleasant with introduction in 2024 year of a new Drilldown interface that simplifies navigation considerably.The old "Explore" interface still remains usable and useful for more complex queries.`}}

{{ tip `I recommend you to invest properly into other forms of monitoring like with Metrics, as they help overview healthy functioning of your application in high performant way, and bring you simplification in investigating problems araised from your next deployments. it is also good idea depending on your application needs investing into Tracing for more deep transperency regarding its performance problems.
Configurations beyond Loki will be overviewed in separate next part articles to keep size of article to a reasonable time of comprehension.` }}

{{ tip `we can build graphical dashboards based on logs alone! It is no efficient in comparison to using metrics, yes it is possible and necessary thing to be used as last resort, or using in low load systems`}}

# Configurating

## Getting server

You need to get somewhere Linux server for deployment (it can be baremetal your own server, or it can VPN rented from some cloud provider). I could recommend Hetzner server, due to the provider being very minimalistic of a high quality and having quite cheap prices. Its Arm64 prices for servers are looking to me like killing feature.

![]({{.StaticRoot}}grafana_loki/hetzner1.png)
![]({{.StaticRoot}}grafana_loki/hetzner2.png)

CAX21 server should be more than enough overkill for our purposes for homelab example purposes. You can squeeze things even into CAX11 if desired, but be mindful preferably to turn on Swap just in case as fallback for ensurancess to handle workload for everything put in at start.

Opentofu(Terraform) code is provided to configure things infrastructure as a code. [See this link for up to date code](https://github.com/darklab8/infra/blob/master/tf/production/node.tf#L5) in case article became outdated.

```c
module "node_darklab_cax21" {
  source     = "../modules/hetzner_server"
  name       = "darklab"
  hardware   = "cax21"
  backups    = true
  ssh_key_id = module.ssh_key.id
  datacenter = "hel1-dc2"
}
```
which utilizes code from this folder https://github.com/darklab8/infra/tree/master/tf/modules/hetzner_server

{{ critical `i highly encourage you to attach Hetzner's firewall to server as it is configured according to this code
https://github.com/darklab8/infra/blob/master/tf/modules/hetzner_server/firewall.tf .
And allow only trafic for 80 and 443 udp and tcp (for our caddy web reverse server), 22(tcp for ssh), icmp for ping ports.

Configured cloud level firewall ensures that in case you forgot something about docker security, you have nice fallback protecting your containers.
That is important with docker, that by default binds applications to 0.0.0.0 when using -p 8000:8000 exposure and it by passes host level firewalls like ufw. Cloud level firewall is your last safe net here in case of human error and missconfigurations.` }}

If u configure the server manually, please create ssh key with `ssh-keygen` command (Usually available right away in linux as long as `git` is installed at least), u can make it available at Windows too if you open Git bash console that becomes available with [installation of git](https://git-scm.com/downloads)

Assuming u created everything correctly, u can make record in your ~/.ssh/config
```
Host homelab
  HostName 65.109.15.108 # replace with IP address shown in hetzner interface
  User root
  IdentityFile ~/.ssh/id_rsa.darklab # replace with name of your SSH key
  IdentitiesOnly yes
```
and connect to it by using `ssh homelab` command. Once u connect and verify `yes`, u will see server insides and be ready for next steps
```
$ ssh homelab
The authenticity of host '65.109.15.108 (65.109.15.108)' can't be established.
ED25519 key fingerprint is SHA256:mQ5+B+9e/1xn3GmRvd0pBnINxtjiLazwT8CMNvI7YcU.
This key is not known by any other names
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '65.109.15.108' (ED25519) to the list of known hosts.
Welcome to Ubuntu 24.04.1 LTS (GNU/Linux 6.8.0-52-generic aarch64)

# bla bla bla, other long text

root@homelab-example:~#
```

## Configuring DNS.

Buy somewnere domain for your server, so that we could have nice address like https://homelab.dd84ai.com for the opening of the web site later with TLS encryption in a named way. We 
Optionally use [free DNS from deSec](https://desec.io/)

Create A record leading leading to public IP of a server

![]({{.StaticRoot}}grafana_loki/dns_record.png)

## Raising docker containers

Once we got server, we can proceed to next step to configuring our monitoring stack. We assume it will be served by Caddy for lets encrypt and reverse proxy

{{ note `We assume you have installed Docker Engine and work from Linux. Instructions may work for WSL2 with Docker Engine or Docker Desktop too, but not guaranteed. With having Docker available locally u will be able to apply instructions from this tutorial without being at server directly.
Instructions for Docker Engine installations can be found here: https://docs.docker.com/engine/install/ubuntu .

If u used Docker app image from hetzner, then Docker is already installed at server.
As last resort u can be just executing this tutorial instructions directly at server, just skip DOCKER_HOST instruction that will be mentioned next`}}

We configure with Docker-compose

{{ note `For convinience of working with some of service running as docker swarm services for easy rotation of their image from CI, we utilize some swarm docker network (overlay) which requires running "docker swarm init" at your server` }}

{{ tip `u can check Opentofu(Terraform) configuration in addition at next address: https://github.com/darklab8/infra/blob/master/tf/modules/docker_stack/monitoring.tf`}}

<div class="code-container">
  <div class="code-tabs">
    <button class="tab-button active" data-lang="yaml">docker-compose.yaml</button>
    <button class="tab-button" data-lang="c">main.tf</button>
  </div>
  <pre class="code-block language-yaml"><code>
{{.MainCompose}}
  </code></pre>
  <pre class="code-block language-c hidden"><code>
{{.MainTerraform}}
  </code></pre>
</div>

Participating configs:

{{ spoiler_start "grafana-datasources.yaml" }}
{{ code_block_start "yaml" }}
{{.GrafanaConfig}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "Dockerfile.grafana" }}
{{ code_block_start "sh" }}
{{.GrafanaDockerfile}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "loki-local-config.yaml" }}
{{ code_block_start "yaml" }}
{{.LokiConfig}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "Dockerfile.loki" }}
{{ code_block_start "sh" }}
{{.LokiDockerfile}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "cfg.logs.alloy" }}
{{ code_block_start "c" }}
{{.AlloyLogsConfig}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

{{ spoiler_start "Dockerfile.alloy.logs" }}
{{ code_block_start "sh" }}
{{.AlloyLogsDockerfile}}
{{ code_block_end "" }}
{{ spoiler_end "" }}

Proceed to apply deployment for raising the monitoring stack
(or use Opentofu(Terraform) to raise all stuff together as modules from ../main.tf)
```sh
git clone https://github.com/darklab8/blog
cd blog/articles/article_detailed/article_20250609_grafana/code_examples

# if docker-compose way
export DOCKER_HOST=ssh://root@homelab
docker ps
docker compose build
export GRAFANA_PASSWORD=INSERT_YOUR_DESIRED_PASSWORD_HERE # replace to your own desired choice
export GRAFANA_DOMAIN=homelab.dd84ai.com # replace to your own desired choice
docker compose up -d

# if opentofu way
tofu init
tofu apply

# after deploy, u need to grant loki proper rights to be persistent
chmod -R a+rw /var/lib/docker/volumes/loki_data
```

Verify everything works correctly.
```
root@homelab-example:/var/lib/docker# docker ps
CONTAINER ID   IMAGE                                   COMMAND                  CREATED          STATUS          PORTS                                                                                NAMES
43248d1628c7   code_examples-alloy-logs                "/bin/alloy run /etc…"   11 minutes ago   Up 11 minutes                                                                                        alloy-logs
f51525f219e7   lucaslorentz/caddy-docker-proxy:2.9.1   "/bin/caddy docker-p…"   11 minutes ago   Up 11 minutes   0.0.0.0:80->80/tcp, :::80->80/tcp, 0.0.0.0:443->443/tcp, :::443->443/tcp, 2019/tcp   caddy
6b2b8c033d6d   code_examples-loki                      "/usr/bin/loki -conf…"   11 minutes ago   Up 11 minutes   3100/tcp                                                                             loki
16ff853a1ba4   code_examples-grafana                   "/run.sh"                11 minutes ago   Up 11 minutes   3000/tcp                                                                             grafana
```
if smth is in status restarting or running suspiciously, check its logs
- by running `docker logs alloy-logs` for example (if container has set explicit easy to use name)
- or `docker logs 43248` (inserting beginning of a container id)
- Alternatively using `docker-compose logs alloy-logs`

## Starting using Grafana

if everything works as intended, u can login into grafana with username `admin` and password u inserted into `GRAFANA_PASSWORD` env var

![]({{.StaticRoot}}grafana_loki/grafana_login.png)

now u can observe now logs for all your running docker containers and navigate them easily by selecting specified log levels

![]({{.StaticRoot}}grafana_loki/loki_drilldown2.png)

or by checking filtering quick to select options. The same is doable at the top of a panel in Labels bar, and Log levels bar.
Filtering by any text simply can be done with menu "Search i nlog lines" and pressing "Include" to specify some search

![]({{.StaticRoot}}grafana_loki/loki_drilldown3.png)

{{ important `Make sure to emit your logs in JSON format! Logging drilldown interface will be automatically recognizing all your json key values as legit labels in a simple way. in Explore and LogQL queries u will need to specify JSON format separately for that to work though.`}}

## Dashboards with Loki

Dashboards using Loki is not very good in performance feature and it has problems to work fast enough for horizontally scaled applications having too many logs, it is more last resort thing to grab essential analytics, and it is good to use feature when u need to use values and their precisions that could jut not fit metrics.
You should configure Mimir/Prometheus with Metrics and compressions through Recording rules for high load applications.
For low workload applications with only single instance of running it is more than enough in performance feature.

To give example onto web like app, we created dummy app example
```sh
export DOCKER_HOST=ssh://root@homelab
docker compose -f docker-compose.app-logs.yaml build
docker compose -f docker-compose.app-logs.yaml up -d
```

now lets create some dashboard using Loki data alone using Loki as data source as flexible "code" mode
![]({{.StaticRoot}}grafana_loki/dashboard_loki1.png)

using new LogQL query with picking possible functions from [Metric Queries](https://grafana.com/docs/loki/latest/query/metric_queries/) page about loki we receive our first dashboard of Max Duration by url pattern over 2m. Notice how we used `unwrap` function to select specific numeric value to be used in formulas, that is important part enabling such things
```
max_over_time({service_name="app-logs"} | json | duration > 0 | url_path!="" | unwrap duration [2m]) by (url_pattern)
```
![]({{.StaticRoot}}grafana_loki/dashboard_loki2.png)

alternatively we could see amount of requests by url pattern in general over 2m
```
sum(count_over_time({service_name="app-logs"} | json | duration > 0 | url_path!="" [2m])) by (url_pattern)
```
![]({{.StaticRoot}}grafana_loki/dashboard_loki3.png)
if you output other information like user IPs, useragents, requests incoming body size, of response body size
you could build then informations of grouping by other parameters, or having shown which URL endpoints use large amount of input or output internet traffic.

That's it for first part of setting up grafana + loki. We will have next article parts, concentrated onto Metrics, Traces and Alerts

