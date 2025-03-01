#!/bin/bash

# 确保脚本出错时退出
set -e

# 检查是否提供了 Docker Hub 用户名
if [ -z "$1" ]; then
    echo "❌  请输入你的 Docker Hub 用户名作为参数！"
    echo "用法: ./build_and_push.sh <your-dockerhub-username>"
    exit 1
fi

# 获取 Docker Hub 用户名
DOCKER_USERNAME=$1
IMAGE_NAME="imghost"
TAG="latest"

# 创建 buildx 构建器（如果不存在）
echo "🚀  正在检查 buildx 构建器..."
if ! docker buildx inspect mybuilder > /dev/null 2>&1; then
    docker buildx create --name mybuilder --use
fi
docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
docker buildx inspect --bootstrap

# 构建并推送多平台镜像 (仅支持 Linux amd64 和 arm64)
echo "🔨  开始构建多平台 Docker 镜像..."
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t $DOCKER_USERNAME/$IMAGE_NAME:$TAG \
  --push .

echo "✅  镜像构建完成并已推送到 Docker Hub：$DOCKER_USERNAME/$IMAGE_NAME:$TAG"
