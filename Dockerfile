FROM golang:1.18.3 as golang
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0 
ENV GOOS=linux 
ENV GOARCH=amd64
ADD . /www
WORKDIR /www
RUN  go mod init work;go mod tidy;go mod vendor;go build -o /www/go-server;chmod 777 /www/go-server

FROM busybox
COPY --from=golang /www/go-server /
ENV ENV local
RUN ls /
EXPOSE 8080
WORKDIR /
ENTRYPOINT ["/go-server"]
