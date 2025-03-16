locals {
  prefix = "${var.common.prefix}-${var.common.env}"
}

# Define ingress ALB
resource "aws_lb" "ingress" {
  name               = "${local.prefix}-alb-ingress"
  internal           = false
  load_balancer_type = "application"
  subnets            = var.network.public_subnet_ids
  security_groups    = [var.network.security_group_for_ingress_alb_id]
  tags = {
    Name = "${local.prefix}-alb-ingress"
  }
}

# Define the listner for ingress ALB
resource "aws_lb_listener" "ingress_prd" {
  load_balancer_arn = aws_lb.ingress.arn
  protocol          = "HTTP"
  port              = "80"
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ingress_blue.arn
  }
}

# Define target group for ingress ALB
resource "aws_lb_target_group" "ingress_blue" {
  name        = "${local.prefix}-tg-backend-blue"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.network.vpc_id

  # TODO: Uncomment the following block to enable health check
  
  health_check {
    protocol            = "HTTP"
    path                = "/"
    port                = "traffic-port"
    healthy_threshold   = 3
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 15
    matcher             = 200
  }
}

# Define the test listner for ingress ALB
resource "aws_lb_listener" "ingress_test" {
  load_balancer_arn = aws_lb.ingress.arn
  protocol          = "HTTP"
  port              = "10080"
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ingress_blue.arn
  }
}

# Define target group Green for ingress ALB
resource "aws_lb_target_group" "ingress_green" {
  name        = "${local.prefix}-tg-backend-green"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.network.vpc_id

  # TODO: Uncomment the following block to enable health check
  health_check {
    protocol            = "HTTP"
    path                = "/"
    port                = "traffic-port"
    healthy_threshold   = 3
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 15
    matcher             = 200
  }
}