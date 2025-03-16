variable "github_actions" {
  type = object({
    account_name = string
    repository   = string
  })
}