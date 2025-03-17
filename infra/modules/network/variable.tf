variable "common" {
  type = object({
    prefix = string
    env    = string
    region = string
  })
}

variable "network" {
  type = object({
    cidr = string
    public_subnets = list(object({
      cidr = string
      az   = string
    }))
  })
}