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

  // Relevant for part 2 article
  tracing = {
    enabled = true
  }
  // Relevant for part 3 article
  metrics = {
    enabled = true
  }
  // Relevant for part 4 article
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
  url  = "https://demo.dd84ai.com/"
  auth = local.grafana_creds
}

// Data sources for all article parts at the same time
module "datasources" {
  source = "./datasources"

  # source = "./infra/tf/modules/grafana_stack/datasources"
  # optionally we can lock ourselves which code to use from external git repo via git source.
  # source = "git@github.com:darklab8/infra.git//tf/modules/grafana_stack/datasources?ref=27d0889348b1b526234d6db7ff60cf2793a772ca"
}
