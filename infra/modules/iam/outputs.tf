output "instance_profile_name" {
  value = aws_iam_instance_profile.ec2_app_profile.name
}

output "ec2_role_name" {
  value = aws_iam_role.ec2_app_role.name
}