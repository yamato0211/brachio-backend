locals {
  prefix = "${var.common.prefix}-${var.common.env}"
  parameters = {
    islocal              = var.backend.islocal
    dynamoendpoint       = var.backend.dynamoendpoint
    cognitosigningkeyurl = var.backend.cognitosigningkeyurl
    poolname             = var.backend.poolname
  }
}

# Define Secrets Manager secret
resource "aws_secretsmanager_secret" "main" {
  name                    = "${local.prefix}/backend"
  description             = "Secret for ${var.common.env} backend application"
  recovery_window_in_days = 0
}

# Define secret values
resource "aws_secretsmanager_secret_version" "main" {
  secret_id     = aws_secretsmanager_secret.main.id
  secret_string = jsonencode(local.parameters)
}