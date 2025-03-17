module "network" {
  source  = "../../modules/network"
  common  = local.common
  network = local.network
}

module "ecr" {
  source = "../../modules/ecr"
  common = local.common
}

module "alb_ingress" {
  source  = "../../modules/alb_ingress"
  common  = local.common
  network = module.network
  domain  = local.domain
}

module "ecs_backend" {
  source          = "../../modules/ecs_backend"
  common          = local.common
  github_actions  = var.github_actions
  network         = module.network
  alb_ingress     = module.alb_ingress
  repository      = module.ecr.backend_repository_arn
  secrets_manager = module.secrets_manager
}

module "frontend" {
  source         = "../../modules/frontend"
  common         = local.common
  github_actions = local.github_actions_for_front
}

module "secrets_manager" {
  source  = "../../modules/secrets_manager"
  common  = local.common
  backend = local.backend_secrets
}