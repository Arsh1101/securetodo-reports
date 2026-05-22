resource "aws_key_pair" "app" {
  key_name   = "${var.project_name}-${var.environment}-managed-key"
  public_key = file(pathexpand(var.public_key_path))

  tags = {
    Name        = "${var.project_name}-${var.environment}-managed-key"
    Project     = var.project_name
    Environment = var.environment
    ManagedBy   = "terraform"
  }
}

module "network" {
  source = "../../modules/network"

  project_name       = var.project_name
  environment        = var.environment
  vpc_cidr           = var.vpc_cidr
  public_subnet_cidr = var.public_subnet_cidr
}

module "storage" {
  source = "../../modules/storage"

  project_name = var.project_name
  environment  = var.environment
}

module "iam" {
  source = "../../modules/iam"

  project_name        = var.project_name
  environment         = var.environment
  reports_bucket_name = module.storage.reports_bucket_name
  reports_bucket_arn  = module.storage.reports_bucket_arn
}

module "compute" {
  source = "../../modules/compute"

  project_name              = var.project_name
  environment               = var.environment
  vpc_id                    = module.network.vpc_id
  public_subnet_id          = module.network.public_subnet_id
  allowed_ssh_cidr          = var.allowed_ssh_cidr
  key_name                  = aws_key_pair.app.key_name
  instance_type             = var.instance_type
  iam_instance_profile_name = module.iam.instance_profile_name

  user_data = file("${path.module}/user_data.sh")
}