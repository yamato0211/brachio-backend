output "backend_repository_arn" {
  value = aws_ecr_repository.backend.arn
}

output "backend_repository_uri" {
  value = aws_ecr_repository.backend.repository_url
}