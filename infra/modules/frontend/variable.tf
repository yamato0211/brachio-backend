variable "own_domain_name" {
  type = object({
    acm_certificate_arn = optional(string)
    aliases             = optional(list(string))
  })
  default = {
    acm_certificate_arn = null
    aliases             = null
  }
}

variable "common" {
  type = object({
    env        = string
    prefix     = string
    account_id = string
  })
}

variable "github_actions" {
  type = object({
    repository   = string
    account_name = string
  })
}