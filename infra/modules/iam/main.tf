locals {
  name_prefix = "${var.project_name}-${var.environment}"
}

resource "aws_iam_role" "ec2_app_role" {
  name = "${local.name_prefix}-ec2-app-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })

  tags = {
    Name        = "${local.name_prefix}-ec2-app-role"
    Project     = var.project_name
    Environment = var.environment
    ManagedBy   = "terraform"
  }
}

resource "aws_iam_policy" "reports_s3_policy" {
  name        = "${local.name_prefix}-reports-s3-policy"
  description = "Least-privilege access for SecureTodo reports bucket."

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "ListReportsBucket"
        Effect = "Allow"
        Action = [
          "s3:ListBucket"
        ]
        Resource = var.reports_bucket_arn
      },
      {
        Sid    = "ReadWriteReports"
        Effect = "Allow"
        Action = [
          "s3:GetObject",
          "s3:PutObject"
        ]
        Resource = "${var.reports_bucket_arn}/reports/*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "reports_s3_attachment" {
  role       = aws_iam_role.ec2_app_role.name
  policy_arn = aws_iam_policy.reports_s3_policy.arn
}

resource "aws_iam_instance_profile" "ec2_app_profile" {
  name = "${local.name_prefix}-ec2-app-profile"
  role = aws_iam_role.ec2_app_role.name
}