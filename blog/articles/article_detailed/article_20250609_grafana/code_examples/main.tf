# Option to raise as Terraform
terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = ">=3.0.2"
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
  source                  = "./infra/tf/modules/docker_stack/monitoring"
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
