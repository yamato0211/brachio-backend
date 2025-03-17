variable "common" {
  type = object({
    prefix = string
    env    = string
    region = string
  })
}
