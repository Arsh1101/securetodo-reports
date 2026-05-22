# Deployment Notes

## Runtime architecture

User traffic reaches nginx on port 80. nginx proxies requests to the Go app on localhost:8080.

```text
User
  ↓
EC2:80
  ↓
nginx
  ↓
localhost:8080
  ↓
SecureTodo Go app
  ↓
SQLite file
  ↓
Local reports or S3 reports
```

## Local container mode

```bash
podman compose up --build
```

## EC2 mode

In EC2 deployment, the app will run as a container and nginx will reverse proxy to it.

## Report storage modes

Local:

```env
REPORT_STORAGE=local
REPORTS_DIR=./reports
```

S3:

```env
REPORT_STORAGE=s3
AWS_REGION=ca-central-1
S3_BUCKET_NAME=<terraform-created-bucket>
S3_REPORTS_PREFIX=reports/
```

## Security notes

- The app container runs as a non-root user.
- The Go app listens on port 8080 internally.
- nginx exposes port 80 publicly.
- SQLite is a local file, not a network database.
- S3 access should use an EC2 IAM role, not hardcoded AWS keys.
- For Fedora/Podman local bind mounts, use `:Z` on mounted volumes.