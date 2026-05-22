#!/usr/bin/env bash

set -euo pipefail

exec > >(tee -a /var/log/securetodo-user-data.log | logger -t securetodo-user-data -s 2>/dev/console) 2>&1

echo "Starting SecureTodo EC2 bootstrap..."

echo "Installing packages..."
dnf update -y
dnf install -y nginx docker openssh-server

echo "Starting SSH..."
systemctl enable sshd
systemctl start sshd

echo "Starting Docker..."
systemctl enable docker
systemctl start docker

for i in {1..30}; do
  if docker info >/dev/null 2>&1; then
    echo "Docker is ready"
    break
  fi

  echo "Waiting for Docker to be ready..."
  sleep 2
done

echo "Writing nginx config..."
cat > /etc/nginx/conf.d/securetodo.conf <<'NGINX'
server {
    listen 80;
    server_name _;

    location / {
        proxy_pass http://127.0.0.1:8080;

        proxy_http_version 1.1;

        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
NGINX

rm -f /etc/nginx/conf.d/default.conf

echo "Testing nginx config..."
nginx -t

echo "Starting nginx..."
systemctl enable nginx
systemctl restart nginx

echo "Bootstrap finished. App container must be loaded and started manually."