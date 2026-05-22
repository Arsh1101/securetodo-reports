#!/usr/bin/env bash

set -euo pipefail

APP_DIR="/opt/securetodo-reports"
APP_NAME="securetodo"
APP_PORT="8080"

dnf update -y
dnf install -y git nginx docker openssh-server

systemctl enable sshd
systemctl start sshd

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

docker info

systemctl enable nginx

rm -rf "$${APP_DIR}"
git clone --branch "${app_repo_branch}" "${app_repo_url}" "$${APP_DIR}"

cd "$${APP_DIR}/app"

docker build -t "$${APP_NAME}:latest" .

docker rm -f "$${APP_NAME}" || true

docker run -d \
  --name "$${APP_NAME}" \
  --restart=always \
  -p 127.0.0.1:$${APP_PORT}:8080 \
  -e APP_ADDR=":8080" \
  -e DB_PATH="./data/securetodo.db" \
  -e REPORT_STORAGE="s3" \
  -e AWS_REGION="${aws_region}" \
  -e S3_BUCKET_NAME="${reports_bucket}" \
  -e S3_REPORTS_PREFIX="reports/" \
  -v securetodo-data:/app/data \
  "$${APP_NAME}:latest"

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

nginx -t

systemctl restart nginx