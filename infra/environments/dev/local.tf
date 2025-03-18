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
    private_subnets = [
      {
        az   = "a"
        cidr = "172.16.10.0/24"
      },
      {
        az   = "c"
        cidr = "172.16.11.0/24"
      }
    ]
  }

  domain = "kurichi.dev"

  github_actions_for_front = {
    account_name = "tosaken1116"
    repository   = "BrachioFront"
  }

  backend_secrets = {
    islocal              = false
    dynamoendpoint       = "https://dynamodb.ap-northeast-1.amazonaws.com"
    cognitosigningkeyurl = "https://cognito-idp.ap-northeast-1.amazonaws.com/ap-northeast-1_LXdcNPdGg/.well-known/jwks.json"
    poolname = "ap-northeast-1_LXdcNPdGg"
  }
}