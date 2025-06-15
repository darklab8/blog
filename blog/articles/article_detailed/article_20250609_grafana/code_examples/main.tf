# Option to raise as Terraform
terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = ">=3.0.2"
    }
    grafana = {
      source = "grafana/grafana"
    }
  }
}

provider "docker" {
  host     = "ssh://homelab"
  ssh_opts = ["-o", "StrictHostKeyChecking=no", "-o", "UserKnownHostsFile=/dev/null", "-i", "~/.ssh/id_rsa.darklab"]
}

module "caddy" {
  source = "./infra/tf/modules/docker_stack/caddy"
}

data "external" "secrets" {
  program = ["pass", "personal/terraform/grafana"]
}

module "monitoring" {
  // Relevant for part 1 article setup and logging
  source = "./infra/tf/modules/docker_stack/monitoring"
  # optionally we can lock ourselves which code to use from external git repo via git source.
  # source = "git@github.com:darklab8/infra.git//tf/modules/docker_stack/monitoring?ref=28407027ebdaba2b48816b63f627c18acd521f46"
  docker_network_caddy_id = module.caddy.network_id
  grafana_password        = data.external.secrets.result["grafana_password"]
  grafana_domain          = "homelab.dd84ai.com"
  logging = {
    enabled = true
  }

  // Relevant for next article parts. Turn off if not needing in part 1.
  tracing = {
    enabled = true
  }
  metrics = {
    enabled = true
  }
  alerts = {
    enabled             = true
    discord_webhook_url = data.external.secrets.result["discord_webhook_url"]
  }
}

locals {
  grafana_password = data.external.secrets.result["grafana_password"]
  grafana_creds    = "admin:${local.grafana_password}"
}


provider "grafana" {
  url  = "https://homelab.dd84ai.com/"
  auth = local.grafana_creds
}

locals {
  loki_uid  = "loki-datasource"
  tempo_uid = "tempo-datasource"
}

resource "grafana_data_source" "loki" {
  type               = "loki"
  name               = "Loki"
  uid                = local.loki_uid
  url                = "http://loki:3100"
  access_mode        = "proxy"
  basic_auth_enabled = false

  json_data_encoded = jsonencode({
    timeout       = 60
    maxLines      = 5000
    tlsSkipVerify = true
    derivedFields = [{
      datasourceUid = local.tempo_uid
      matcherRegex  = "trace_id"
      name          = "trace_id"
      url : "$${__value.raw}"
      urlDisplayLabel : "trace_id"
      matcherType = "label"
    }]
  })
  depends_on = [
    module.monitoring
  ]
}