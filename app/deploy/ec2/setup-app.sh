#!/usr/bin/env bash

set -euo pipefail

APP_DIR="/opt/securetodo-reports"

sudo mkdir -p "$APP_DIR"
sudo chown "$USER":"$USER" "$APP_DIR"

echo "Application directory prepared: $APP_DIR"
echo "Next phase: Terraform user_data will automate installing Podman, nginx, and starting the app."