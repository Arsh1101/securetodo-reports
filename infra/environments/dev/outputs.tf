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
  value = "ssh -i <path-to-private-key> ec2-user@${module.compute.public_ip}"
}

output "reports_bucket_name" {
  value = module.storage.reports_bucket_name
}

output "ec2_role_name" {
  value = module.iam.ec2_role_name
}