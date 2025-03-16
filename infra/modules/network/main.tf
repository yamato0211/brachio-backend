locals {
  prefix = "${var.common.prefix}-${var.common.env}"
}

# VPC
resource "aws_vpc" "main" {
  cidr_block = var.network.cidr

  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = {
    Name = "${local.prefix}-vpc"
  }
}

# Public Subnets
resource "aws_subnet" "public_ingress" {
  for_each          = { for i, s in var.network.public_subnets : i => s }
  vpc_id            = aws_vpc.main.id
  availability_zone = "${var.common.region}${each.value.az}"
  cidr_block        = each.value.cidr
  tags = {
    Name = "${local.prefix}-subnet-public-1${each.value.az}"
  }
}

# Internet Gateway
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "${local.prefix}-igw"
  }
}

# Route Table
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "${local.prefix}-rtb-public"
  }
}

# Route
resource "aws_route" "public" {
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.main.id
  route_table_id         = aws_route_table.public.id
}

# Route Table Association
resource "aws_route_table_association" "public" {
  for_each       = aws_subnet.public_ingress
  subnet_id      = each.value.id
  route_table_id = aws_route_table.public.id
}

# Define security group for management
resource "aws_security_group" "management" {
  name   = "${local.prefix}-sg-management"
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "${local.prefix}-sg-management"
  }
}

resource "aws_vpc_security_group_egress_rule" "management" {
  security_group_id = aws_security_group.management.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

# Define security group for ingress ALB
resource "aws_security_group" "ingress_alb" {
  name   = "${local.prefix}-sg-ingress-alb"
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "${local.prefix}-sg-ingress-alb"
  }
}

resource "aws_vpc_security_group_ingress_rule" "ingress_alb" {
  security_group_id = aws_security_group.ingress_alb.id
  ip_protocol = "tcp"
  from_port = 80
  to_port = 80
  cidr_ipv4 = "0.0.0.0/0"
} 

resource "aws_vpc_security_group_egress_rule" "ingress_alb" {
  security_group_id = aws_security_group.ingress_alb.id
  ip_protocol = "-1"
  cidr_ipv4 = "0.0.0.0/0"
}

# Define security group for backend application
resource "aws_security_group" "backend" {
  name = "${local.prefix}-sg-backend"
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "${local.prefix}-sg-backend"
  }
}

resource "aws_vpc_security_group_ingress_rule" "backend" {
  security_group_id = aws_security_group.backend.id
  ip_protocol = "tcp"
  from_port = 80
  to_port = 80
  referenced_security_group_id = aws_security_group.ingress_alb.id
} 

resource "aws_vpc_security_group_egress_rule" "backend" {
  security_group_id = aws_security_group.backend.id
  ip_protocol = "-1"
  cidr_ipv4 = "0.0.0.0/0"
}
