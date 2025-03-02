# 使用最新的 Golang 1.23 作为构建环境
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . .

# 编译二进制文件（静态编译以减少依赖）
RUN CGO_ENABLED=0 GOOS=linux go build -o objectStorage

# 使用更小的基础镜像运行
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 复制编译好的二进制文件
COPY --from=builder /app/objectStorage /app/

# 运行时使用的端口（如果你的应用监听特定端口）
EXPOSE 8080

# 运行应用
CMD ["./objectStorage"]
