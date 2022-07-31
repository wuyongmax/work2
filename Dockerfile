FROM golang:1.18.3 as golang
ENV GO111MODULE=off
ENV GOPROXY=https://goproxy.cn,direct
ADD . /www
WORKDIR /www
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -installsuffix cgo -o /www/go-server main.go;chmod 777 /www/go-server
FROM busybox
COPY --from=golang /www/go-server /
ENV ENV local
RUN ls /
EXPOSE 8080
WORKDIR /
ENTRYPOINT ["/go-server"]
