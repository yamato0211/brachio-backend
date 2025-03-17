variable "common" {
  type = object({
    prefix     = string
    env        = string
    region     = string
    account_id = string
  })
}

variable "github_actions" {
  type = object({
    account_name = string
    repository   = string
  })
}

variable "network" {
  type = object({
    vpc_id                                  = string
    public_subnet_ids                       = list(string)
    security_group_for_backend_container_id = string
  })
}

variable "alb_ingress" {
  type = object({
    alb_listener_ingress_prd_arn        = string
    alb_target_group_ingress_blue_name  = string
    alb_target_group_ingress_blue_arn   = string
    alb_listener_ingress_test_arn       = string
    alb_target_group_ingress_green_name = string
    alb_target_group_ingress_green_arn  = string
  })
}

# variable "secrets_manager" {
#   type = object({
#     secret_for_db_arn = string
#   })
# }

variable "repository" {
  type = string
}

variable "secrets_manager" {
  type = object({
    secret_for_backend_arn = string
  })
}