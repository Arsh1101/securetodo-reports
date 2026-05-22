#!/usr/bin/env bash

set -euo pipefail

dnf update -y
dnf install -y nginx

systemctl enable nginx
systemctl start nginx

cat > /usr/share/nginx/html/index.html <<'HTML'
<!doctype html>
<html>
<head>
  <title>SecureTodo Terraform Demo</title>
</head>
<body>
  <h1>SecureTodo Terraform Demo</h1>
  <p>EC2 and nginx were provisioned by Terraform.</p>
  <p>The Go app container will be added in the next phase.</p>
</body>
</html>
HTML