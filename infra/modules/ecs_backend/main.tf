locals {
  prefix = "${var.common.prefix}-${var.common.env}"
}

resource "aws_s3_bucket" "backend" {
  bucket = "${local.prefix}-backend-bucket"
}

resource "aws_iam_role" "task_role" {
  name = "${local.prefix}-task-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_policy" "task_policy" {
  name = "${local.prefix}-task-policy"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "S3ObjectAccess",
        Effect = "Allow",
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject"
        ],
        Resource = "arn:aws:s3:::${aws_s3_bucket.backend.arn}/*"
      },
      {
        Sid      = "LambdaFullAccess",
        Effect   = "Allow",
        Action   = "lambda:*",
        Resource = "*"
      },
      {
        Sid    = "CognitoAccess",
        Effect = "Allow",
        Action = [
          "cognito-idp:*",
        ],
        Resource = "*"
      },
      {
        Sid    = "DynamoDBAccess",
        Effect = "Allow",
        Action = [
          "dynamodb:*"
        ],
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "task_role_attachment" {
  role       = aws_iam_role.task_role.name
  policy_arn = aws_iam_policy.task_policy.arn
}

# Define ECS task definition
resource "aws_ecs_task_definition" "backend" {
  family                   = "${local.prefix}-backend-def"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 512
  memory                   = 1024
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.task_execution_role.arn
  task_role_arn            = aws_iam_role.task_role.arn
  container_definitions = jsonencode([
    {
      name      = "server-app"
      image     = "${var.common.account_id}.dkr.ecr.ap-northeast-1.amazonaws.com/brachio-dev-backend:sha-788ef28"
      cpu       = 256
      memory    = 512
      essential = true

      secrets = [
        {
          name      = "IS_LOCAL"
          valueFrom = "${var.secrets_manager.secret_for_backend_arn}:islocal::"
        },
        {
          name      = "DYNAMO_ENDPOINT"
          valueFrom = "${var.secrets_manager.secret_for_backend_arn}:dynamoendpoint::"
        },
        {
          name      = "SIGNING_KEY_URL"
          valueFrom = "${var.secrets_manager.secret_for_backend_arn}:cognitosigningkeyurl::"
        },
        {
          name      = "USER_POOL_NAME"
          valueFrom = "${var.secrets_manager.secret_for_backend_arn}:poolname::"
        }
      ]

      portMappings = [{ containerPort = 8080 }]
      "readonlyRootFilesystem" : true
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-region : "ap-northeast-1"
          awslogs-group : aws_cloudwatch_log_group.backend.name
          awslogs-stream-prefix : "ecs"
        }
      }
    }
  ])
}

# Define ECS cluster
resource "aws_ecs_cluster" "backend" {
  name = "${local.prefix}-backend-cluster"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

# Define ECS service
resource "aws_ecs_service" "backend" {
  name                               = "${local.prefix}-ecs-backend-service"
  cluster                            = aws_ecs_cluster.backend.arn
  task_definition                    = aws_ecs_task_definition.backend.arn
  launch_type                        = "FARGATE"
  platform_version                   = "1.4.0"
  scheduling_strategy                = "REPLICA"
  desired_count                      = 1
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent         = 200
  deployment_controller {
    type = "CODE_DEPLOY"
  }
  enable_ecs_managed_tags = true
  network_configuration {
    subnets = var.network.public_subnet_ids
    security_groups = [
      var.network.security_group_for_backend_container_id
    ]
    assign_public_ip = true
  }
  health_check_grace_period_seconds = 120
  load_balancer {
    target_group_arn = var.alb_ingress.alb_target_group_ingress_blue_arn
    container_name   = "server-app"
    container_port   = 8080
  }
  lifecycle {
    ignore_changes = [
      task_definition,
      load_balancer,
      network_configuration
    ]
  }
}

# Define ECS task execution role
resource "aws_iam_role" "task_execution_role" {
  name               = "${local.prefix}-task-execution-role"
  assume_role_policy = data.aws_iam_policy_document.trust_policy_for_task_execution_role.json
}

data "aws_iam_policy_document" "trust_policy_for_task_execution_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_policy" "policy_for_access_to_secrets_manager" {
  name   = "${local.prefix}-GettingSecretsPolicy-backend"
  policy = data.aws_iam_policy_document.policy_for_access_to_secrets_manager.json
}

data "aws_iam_policy_document" "policy_for_access_to_secrets_manager" {
  statement {
    effect    = "Allow"
    resources = ["*"]
    actions = [
      "secretsmanager:GetSecretValue",
    ]
  }
}

resource "aws_iam_role_policy_attachment" "task_execution_role" {
  for_each = {
    ecs            = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
    secretsmanager = aws_iam_policy.policy_for_access_to_secrets_manager.arn
  }
  role       = aws_iam_role.task_execution_role.name
  policy_arn = each.value
}

# Define CodeDeploy application
resource "aws_codedeploy_app" "backend" {
  compute_platform = "ECS"
  name             = "${local.prefix}-backend-app"
}

# Define CodeDeploy deployment group
resource "aws_codedeploy_deployment_group" "backend" {
  app_name               = aws_codedeploy_app.backend.name
  deployment_config_name = "CodeDeployDefault.ECSAllAtOnce"
  deployment_group_name  = "${local.prefix}-ecs-backend-deployment-group"
  service_role_arn       = aws_iam_role.codedeploy.arn
  deployment_style {
    deployment_option = "WITH_TRAFFIC_CONTROL"
    deployment_type   = "BLUE_GREEN"
  }
  ecs_service {
    cluster_name = aws_ecs_cluster.backend.name
    service_name = aws_ecs_service.backend.name
  }
  load_balancer_info {
    target_group_pair_info {
      prod_traffic_route {
        listener_arns = [var.alb_ingress.alb_listener_ingress_prd_arn]
      }
      test_traffic_route {
        listener_arns = [var.alb_ingress.alb_listener_ingress_test_arn]
      }
      target_group {
        name = var.alb_ingress.alb_target_group_ingress_blue_name
      }
      target_group {
        name = var.alb_ingress.alb_target_group_ingress_green_name
      }
    }
  }
  blue_green_deployment_config {
    deployment_ready_option {
      action_on_timeout    = "STOP_DEPLOYMENT"
      wait_time_in_minutes = 10
    }
    terminate_blue_instances_on_deployment_success {
      action                           = "TERMINATE"
      termination_wait_time_in_minutes = 60
    }
  }
}

# Define IAM role for CodeDeploy
resource "aws_iam_role" "codedeploy" {
  name               = "${local.prefix}-role-for-codedeploy"
  assume_role_policy = data.aws_iam_policy_document.trust_policy_for_codedeploy.json
}

data "aws_iam_policy_document" "trust_policy_for_codedeploy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["codedeploy.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role_policy_attachment" "codedeploy" {
  for_each = {
    "ecs" = "arn:aws:iam::aws:policy/AWSCodeDeployRoleForECS"
  }
  role       = aws_iam_role.codedeploy.name
  policy_arn = each.value
}

# Define CloudWatch log group
resource "aws_cloudwatch_log_group" "backend" {
  name              = "/ecs/${local.prefix}-backend"
  retention_in_days = 7

  tags = {
    Name = "/ecs/${local.prefix}-backend"
  }
}

# Define IAM role for Github Actions
resource "aws_iam_role" "github_actions" {
  name               = "${local.prefix}-role-for-github-actions"
  assume_role_policy = data.aws_iam_policy_document.trust_policy_for_github_actions.json
}

data "aws_iam_policy_document" "trust_policy_for_github_actions" {
  statement {
    effect = "Allow"
    principals {
      type        = "Federated"
      identifiers = ["arn:aws:iam::${var.common.account_id}:oidc-provider/token.actions.githubusercontent.com"]
    }
    actions = ["sts:AssumeRoleWithWebIdentity"]
    condition {
      test     = "StringLike"
      variable = "token.actions.githubusercontent.com:sub"
      values   = ["repo:${var.github_actions.account_name}/${var.github_actions.repository}:*"]
    }
  }
}

resource "aws_iam_policy" "policy_for_github_actions" {
  name   = "${local.prefix}-policy-for-github-actions"
  policy = data.aws_iam_policy_document.policy_for_github_actions.json
}

data "aws_iam_policy_document" "policy_for_github_actions" {
  statement {
    sid    = "GetAuthorizationToken"
    effect = "Allow"
    actions = [
      "ecr:GetAuthorizationToken"
    ]
    resources = ["*"]
  }
  statement {
    sid    = "PushImageOnly"
    effect = "Allow"
    actions = [
      "ecr:BatchCheckLayerAvailability",
      "ecr:BatchGetImage",
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
      "ecr:PutImage"
    ]
    resources = [var.repository]
  }
  statement {
    sid    = "RegisterTaskDefinition"
    effect = "Allow"
    actions = [
      "ecs:RegisterTaskDefinition",
      "ecs:DescribeTaskDefinition"
    ]
    resources = ["*"]
  }
  statement {
    sid    = "UpdateService"
    effect = "Allow"
    actions = [
      "ecs:UpdateServicePrimaryTaskSet",
      "ecs:DescribeServices",
      "ecs:UpdateService"
    ]
    resources = [aws_ecs_service.backend.id]
  }
  statement {
    sid    = "PassRole"
    effect = "Allow"
    actions = [
      "iam:PassRole"
    ]
    resources = [
      aws_iam_role.task_execution_role.arn,
      aws_iam_role.task_role.arn
    ]
    condition {
      test     = "StringLike"
      variable = "iam:PassedToService"
      values   = ["ecs-tasks.amazonaws.com"]
    }
  }
  statement {
    sid    = "DeployService"
    effect = "Allow"
    actions = [
      "codedeploy:CreateDeployment",
      "codedeploy:GetDeployment",
      "codedeploy:GetDeploymentConfig",
      "codedeploy:GetDeploymentGroup",
      "codedeploy:RegisterApplicationRevision"
    ]
    resources = [
      aws_codedeploy_app.backend.arn,
      aws_codedeploy_deployment_group.backend.arn,
      "arn:aws:codedeploy:${var.common.region}:${var.common.account_id}:deploymentconfig:*"
    ]
  }
}

resource "aws_iam_role_policy_attachment" "github_actions" {
  for_each = {
    iam    = "arn:aws:iam::aws:policy/IAMReadOnlyAccess",
    github = aws_iam_policy.policy_for_github_actions.arn
  }
  role       = aws_iam_role.github_actions.name
  policy_arn = each.value
}

resource "aws_dynamodb_table" "card_table" {
  name     = "MasterCards"
  hash_key = "MasterCardId"

  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "MasterCardId"
    type = "S"
  }

  tags = {
    Name        = "${local.prefix}-dynamodb-master-card-table"
    Environment = var.common.env
  }
}

resource "aws_dynamodb_table" "user_table" {
  name     = "Users"
  hash_key = "UserId"

  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "UserId"
    type = "S"
  }

  tags = {
    Name        = "${local.prefix}-dynamodb-user-table"
    Environment = var.common.env
  }
}

resource "aws_dynamodb_table" "item_table" {
  name     = "MasterItems"
  hash_key = "MasterItemId"

  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "MasterItemId"
    type = "S"
  }

  tags = {
    Name        = "${local.prefix}-dynamodb-master-item-table"
    Environment = var.common.env
  }
}

resource "aws_dynamodb_table" "deck_table" {
  name     = "Decks"
  hash_key = "DeckId"

  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "DeckId"
    type = "S"
  }

  attribute {
    name = "UserId"
    type = "S"
  }

  global_secondary_index {
    name            = "UserIdIndex"
    hash_key        = "UserId"
    projection_type = "ALL"
  }

  tags = {
    Name        = "${local.prefix}-dynamodb-deck-table"
    Environment = var.common.env
  }
}

resource "aws_dynamodb_table" "present_table" {
  name      = "Presents"
  hash_key  = "PresentId"
  range_key = "Time"

  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "PresentId"
    type = "S"
  }

  attribute {
    name = "Time"
    type = "N"
  }

  tags = {
    Name        = "${local.prefix}-dynamodb-present-table"
    Environment = var.common.env
  }
}

resource "aws_iam_policy" "dynamodb_policy" {
  name = "${local.prefix}-dynamodb-policy"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "DynamoDBFullAccess",
        Effect = "Allow",
        Action = "dynamodb:*",
        Resource = [
          aws_dynamodb_table.card_table.arn,
          aws_dynamodb_table.user_table.arn,
          aws_dynamodb_table.item_table.arn,
          aws_dynamodb_table.deck_table.arn,
          aws_dynamodb_table.present_table.arn
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "dynamodb_policy_attachment" {
  role       = aws_iam_role.task_role.name
  policy_arn = aws_iam_policy.dynamodb_policy.arn
}
