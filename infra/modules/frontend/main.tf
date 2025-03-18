locals {
  prefix = "${var.common.prefix}-${var.common.env}"
}

resource "aws_s3_bucket_policy" "s3_policy" {
  bucket = aws_s3_bucket.default.id
  policy = jsonencode({
    Version = "2008-10-17"
    Id      = "PolicyForCloudFrontPrivateContent"
    Statement = [
      {
        Sid    = "AllowCloudFrontServicePrincipal"
        Effect = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action   = "s3:GetObject"
        Resource = "arn:aws:s3:::${aws_s3_bucket.default.bucket}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = aws_cloudfront_distribution.cf.arn
          }
        }
      }
    ]
  })
}

resource "aws_s3_bucket" "default" {
  bucket        = "${local.prefix}-frontend"
  force_destroy = true
}

resource "aws_cloudfront_origin_access_control" "default" {
  name                              = "${local.prefix}-frontend-oac"
  description                       = "S3のオブジェクトをインターネットへ配信するためのOAC"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_cloudfront_distribution" "cf" {
  comment = aws_s3_bucket.default.bucket

  enabled          = "true"
  http_version     = "http2"
  is_ipv6_enabled  = "true"
  price_class      = "PriceClass_All"
  retain_on_delete = "false"

  default_root_object = "index.html"

  aliases = ["pokepoke.${var.domain}"]
  origin {
    domain_name              = aws_s3_bucket.default.bucket_regional_domain_name
    origin_id                = aws_s3_bucket.default.bucket
    origin_access_control_id = aws_cloudfront_origin_access_control.default.id
    connection_attempts      = "3"
    connection_timeout       = "10"
  }

  custom_error_response {
    error_caching_min_ttl = 0
    error_code            = 403
    response_code         = 200
    response_page_path    = "/index.html"
  }

  custom_error_response {
    error_caching_min_ttl = 0
    error_code            = 404
    response_code         = 200
    response_page_path    = "/index.html"
  }

  viewer_certificate {
    # 独自ドメインのACMを使用しない場合、デフォルトの証明書を使用する
    cloudfront_default_certificate = false
    # cloudfront_default_certificate = aws_acm_certificate.main.arn

    # 独自ドメインのACMを使用する場合、証明書を指定する
    # acm_certificate_arn = var.own_domain_name.acm_certificate_arn == null ? null : var.own_domain_name.acm_certificate_arn

    acm_certificate_arn = aws_acm_certificate.front.arn

    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1"
  }

  default_cache_behavior {
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    compress               = "true"
    default_ttl            = "60"
    max_ttl                = "60"
    min_ttl                = "60"
    smooth_streaming       = "false"
    target_origin_id       = aws_s3_bucket.default.id
    viewer_protocol_policy = "redirect-to-https"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }
  }

  restrictions {
    geo_restriction { // 地理的制限
      restriction_type = "none"
    }
  }
}

resource "aws_iam_role" "github_actions" {
  name               = "${local.prefix}-role-for-frontend-github-actions"
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
  name   = "${local.prefix}-policy-for-frontend-github-actions"
  policy = data.aws_iam_policy_document.policy_for_github_actions.json
}

data "aws_iam_policy_document" "policy_for_github_actions" {
  # S3 バケット自体の一覧取得を許可
  statement {
    sid    = "S3ListBucket"
    effect = "Allow"
    actions = [
      "s3:ListBucket"
    ]
    resources = ["arn:aws:s3:::${aws_s3_bucket.default.bucket}"]
  }
  # S3 バケット内のオブジェクトへのアクセス（アップロード・削除など）を許可
  statement {
    sid    = "S3ObjectActions"
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject",
      "s3:PutObjectAcl"
    ]
    resources = ["arn:aws:s3:::${aws_s3_bucket.default.bucket}/*"]
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

// Host Zone作成
resource "aws_route53_zone" "front" {
  name = "pokepoke.${var.domain}"
}

resource "aws_route53_record" "front" {
  zone_id = aws_route53_zone.front.id
  name    = "pokepoke.${var.domain}"
  type    = "A"
  alias {
    name                   = aws_cloudfront_distribution.cf.domain_name
    zone_id                = aws_cloudfront_distribution.cf.hosted_zone_id
    evaluate_target_health = true
  }
}

// ACM
resource "aws_acm_certificate" "front" {
  domain_name       = "pokepoke.${var.domain}"
  validation_method = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "front" {
  certificate_arn         = aws_acm_certificate.front.arn
  validation_record_fqdns = [for record in aws_route53_record.cert : record.fqdn]
}

resource "aws_route53_record" "cert" {
  for_each = {
    for dvo in aws_acm_certificate.front.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      record = dvo.resource_record_value
    }
  }
  zone_id = aws_route53_zone.front.id
  name    = each.value.name
  type    = each.value.type
  ttl     = 600
  records = [each.value.record]
}

