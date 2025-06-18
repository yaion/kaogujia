# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.21-alpine

# 设置工作目录
WORKDIR /app

# 拷贝项目源码到容器中
COPY . .

# 下载依赖
RUN go mod download

# 构建应用，输出可执行文件
RUN go build -o /go-app

# 暴露应用运行的端口（根据实际需求修改）
EXPOSE 8080

# 容器启动时运行的命令
CMD ["/kaogujia"]