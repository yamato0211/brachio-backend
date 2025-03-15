data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

locals {
  common = {
    prefix     = "brachio"
    env        = "dev"
    region     = data.aws_region.current.name
    account_id = data.aws_caller_identity.current.account_id
  }

  network = {
    cidr = "172.16.0.0/16"
    public_subnets = [
      {
        az   = "a"
        cidr = "172.16.0.0/24"
      },
      {
        az   = "c"
        cidr = "172.16.1.0/24"
      }
    ]
  }
}