output "s3_bucket_id" { value = aws_s3_bucket.default.id }
output "s3_bucket_regional_domain_name" { value = aws_s3_bucket.default.bucket_regional_domain_name }
output "cloudfront_domain_name" { value = aws_cloudfront_distribution.cf.domain_name }
output "cloudfront_aliases" { value = aws_cloudfront_distribution.cf.aliases }
output "cloudfront_arn" { value = aws_cloudfront_distribution.cf.arn }
output "cloudfront_hosted_zone_id" { value = aws_cloudfront_distribution.cf.hosted_zone_id }