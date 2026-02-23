#!/bin/bash

echo "--- Tidy Go modules ---"
go mod tidy

#启动docker
echo "--- Starting Docker Compose ---"
docker-compose up -d

# 等待中间件启动完成（根据需要调整秒数）
echo "Waiting for middleware to be ready..."
sleep 5

SERVICES=("gateway" "product" "user" "order" "cart" "payment" "stock" "coupon")

for SVC in "${SERVICES[@]}"; do
    echo "Starting service: $SVC"
    # 后台运行 go run，并将日志重定向到 logs 目录
    mkdir -p logs
    nohup go run cmd/$SVC/main.go > logs/info.log 2>&1 &
done

echo "--- All services started ---"
echo "Check logs/ directory for details."