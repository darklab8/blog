package article_20250609_grafana

import (
	_ "embed"
)

type Vars struct {
	StaticRoot string
	SiteRoot   string

	MainCompose   string
	MainTerraform string

	GrafanaConfig     string
	GrafanaDockerfile string

	LokiConfig     string
	LokiDockerfile string

	AlloyLogsConfig     string
	AlloyLogsDockerfile string

	AppLogsDashboard string

	HetznerFirewall string
}

//go:embed code_examples/docker-compose.yaml
var MainCompose string

//go:embed code_examples/main.tf
var MainTerraform string

//go:embed code_examples/grafana-datasources.yaml
var GrafanaConfig string

//go:embed code_examples/Dockerfile.grafana
var GrafanaDockerfile string

//go:embed code_examples/infra/tf/modules/docker_stack/monitoring/loki-local-config.yaml
var LokiConfig string

//go:embed code_examples/Dockerfile.loki
var LokiDockerfile string

//go:embed code_examples/infra/tf/modules/docker_stack/monitoring/cfg.logs.alloy
var AlloyLogsConfig string

//go:embed code_examples/Dockerfile.alloy.logs
var AlloyLogsDockerfile string

//go:embed code_examples/dashboard_app_logs.json
var AppLogsDashboard string

//go:embed code_examples/infra/tf/modules/hetzner_server/firewall.tf
var HetznerFirewall string
