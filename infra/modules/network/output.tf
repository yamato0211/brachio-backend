########################
# VPC
########################

output "vpc_id" {
  value = aws_vpc.main.id
}

########################
# Subnet
########################

output "public_subnet_ids" {
  value = values(aws_subnet.public_ingress)[*].id
}

########################
# Security Group
########################

output "security_group_for_management_id" {
  value = aws_security_group.management.id
}

output "security_group_for_ingress_alb_id" {
  value = aws_security_group.ingress_alb.id
}

output "security_group_for_backend_container_id" {
  value = aws_security_group.backend.id
}
