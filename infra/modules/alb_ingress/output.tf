output "alb_listener_ingress_prd_arn" {
  value = aws_lb_listener.ingress_prd.arn
}

output "alb_target_group_ingress_blue_name" {
  value = aws_lb_target_group.ingress_blue.name
}

output "alb_target_group_ingress_blue_arn" {
  value = aws_lb_target_group.ingress_blue.arn
}

output "alb_listener_ingress_test_arn" {
  value = aws_lb_listener.ingress_test.arn
}

output "alb_target_group_ingress_green_name" {
  value = aws_lb_target_group.ingress_green.name
}

output "alb_target_group_ingress_green_arn" {
  value = aws_lb_target_group.ingress_green.arn
}