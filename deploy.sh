#!/bin/bash

# Kiểm tra xem API key có tồn tại không
if [ -z "$RENDER_API_KEY" ]; then
  echo "RENDER_API_KEY is required"
  exit 1
fi

# Đăng nhập vào Render và triển khai ứng dụng
echo "Deploying to Render..."
curl -X POST https://api.render.com/v1/services/email-service/deployments \
  -H "Authorization: Bearer $RENDER_API_KEY"