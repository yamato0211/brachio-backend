locals {
  prefix = "${var.common.prefix}-${var.common.env}"
}

# Define ECR repository for backend app
resource "aws_ecr_repository" "backend" {
  name                 = "${local.prefix}-backend"
  image_tag_mutability = "MUTABLE"
  force_delete         = true
  image_scanning_configuration {
    scan_on_push = false
  }
  encryption_configuration {
    encryption_type = "KMS"
  }

  tags = {
    Name = "${local.prefix}-ecr-backend"
  }
}