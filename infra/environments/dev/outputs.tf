output "app_url" {
  value = "http://${module.compute.public_ip}"
}

output "ec2_public_ip" {
  value = module.compute.public_ip
}

output "ec2_public_dns" {
  value = module.compute.public_dns
}

output "ssh_command" {
  value = "ssh -i ~/.ssh/securetodo-dev-managed-key ec2-user@${module.compute.public_ip}"
}

output "reports_bucket_name" {
  value = module.storage.reports_bucket_name
}

output "ec2_role_name" {
  value = module.iam.ec2_role_name
}

output "security_group_id" {
  value = module.compute.security_group_id
}

output "instance_id" {
  value = module.compute.instance_id
}

output "ecr_repository_name" {
  value = module.ecr.repository_name
}

output "ecr_repository_url" {
  value = module.ecr.repository_url
}