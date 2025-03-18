terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.91.0"
    }
  }

  required_version = ">= 1.2.0"

  backend "s3" {
    bucket = "brachio-terraform-state"
    key    = "terraform-dev.tfstate"
    region = "ap-northeast-1"
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

provider "aws" {
  region = "us-east-1"
  alias  = "us-east-1"
}