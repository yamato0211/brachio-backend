variable "common" {
  type = object({
    prefix = string
    env    = string
    region = string
  })
}

variable "network" {
  type = object({
    vpc_id                            = string
    public_subnet_ids                 = list(string)
    security_group_for_ingress_alb_id = string
  })
}

variable "domain" {
  type = string
}