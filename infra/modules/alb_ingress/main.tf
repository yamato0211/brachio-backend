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
  protocol          = "HTTPS"
  port              = "443"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = aws_acm_certificate.main.arn
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ingress_blue.arn
  }
}

resource "aws_lb_listener" "ingress_prd_http" {
  load_balancer_arn = aws_lb.ingress.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"
    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}
# Define target group for ingress ALB
resource "aws_lb_target_group" "ingress_blue" {
  name        = "${local.prefix}-tg-backend-blue"
  port        = 8080
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
  protocol          = "HTTPS"
  port              = "4403"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = aws_acm_certificate.main.arn
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.ingress_blue.arn
  }
}

resource "aws_lb_listener" "ingress_test_http" {
  load_balancer_arn = aws_lb.ingress.arn
  port              = 10080
  protocol          = "HTTP"
  default_action {
    type = "redirect"
    redirect {
      port        = "4403"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

# Define target group Green for ingress ALB
resource "aws_lb_target_group" "ingress_green" {
  name        = "${local.prefix}-tg-backend-green"
  port        = 8080
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

// Host Zone作成
resource "aws_route53_zone" "main" {
  name = "brachio.${var.domain}"
}

resource "aws_route53_record" "prod" {
  zone_id = aws_route53_zone.main.id
  name    = "brachio.${var.domain}"
  type    = "A"
  alias {
    name                   = aws_lb.ingress.dns_name
    zone_id                = aws_lb.ingress.zone_id
    evaluate_target_health = true
  }
}

// ACM
resource "aws_acm_certificate" "main" {
  domain_name       = "brachio.${var.domain}"
  validation_method = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "main" {
  certificate_arn         = aws_acm_certificate.main.arn
  validation_record_fqdns = [for record in aws_route53_record.cert : record.fqdn]
}

resource "aws_route53_record" "cert" {
  for_each = {
    for dvo in aws_acm_certificate.main.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      record = dvo.resource_record_value
    }
  }
  zone_id = aws_route53_zone.main.id
  name    = each.value.name
  type    = each.value.type
  ttl     = 600
  records = [each.value.record]
}

