variable "common" {
  type = object({
    env    = string
    prefix = string
  })
}

variable "backend" {
  type = object({
    islocal              = bool
    dynamoendpoint       = string
    cognitosigningkeyurl = string
  })
}