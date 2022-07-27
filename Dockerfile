FROM golang:1.17-alpine as golang
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ADD . /www
WORKDIR /www
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod init work;go mod tidy;go mod vendor;go build -o /www/go-server;chmod 777 /www/go-server
ENTRYPOINT ["/www/go-server"]
