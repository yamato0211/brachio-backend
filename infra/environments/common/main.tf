module "oidc-github_example_complete" {
  source              = "unfunco/oidc-github/aws//examples/complete"
  version             = "1.8.1"
  github_repositories = ["yamato0211/brachio-backend"]
  region              = "ap-northeast-1"
  iam_role_name       = "github-actions-for-terraform"
}