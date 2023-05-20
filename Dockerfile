FROM golang:1.18 as builder
WORKDIR /workplace/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build  -v -mod=vendor -o app src/main.go

# alpine
FROM registry.cn-shanghai.aliyuncs.com/pub_space/alpine:base

WORKDIR /workplace/
COPY --from=builder /workplace/app .
COPY --from=builder /workplace/src/config/config.toml config.toml

EXPOSE 80
CMD ["/workplace/app","-p=80"]