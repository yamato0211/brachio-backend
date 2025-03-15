terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"

  backend "s3" {
    bucket = "brachio-terraform-state"
    key = "terraform.tfstate"
    region = "ap-northeast-1"
  }
}

provider "aws" {
  region  = "ap-northeast-1"
}

resource "aws_lambda_function" "my_hello_world" {
  runtime = "nodejs12.x"
  handler = "exports.test"
  image_uri = "test"
  function_name = "test"
  role = "arn:aws:ec2:us-east-1:123123123123:instance/i-1231231231"

  memory_size = 512
  tags = {
    Environment = "dev"
  }
}
