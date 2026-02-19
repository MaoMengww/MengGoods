FROM golang:1.25.1 AS builder

ARG SERVICE_NAME

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go mod tidy && go build -o server ./cmd/${SERVICE_NAME}/main.go

FROM ubuntu:latest

WORKDIR /app

#拷贝二进制文件
COPY --from=builder /app/server .

COPY --from=builder /app/config ./config

RUN apt-get update && apt-get install -y tzdata && rm -rf /var/lib/apt/lists/*
ENV TZ=Asia/Shanghai

EXPOSE 8080

CMD [ "/app/server" ]


