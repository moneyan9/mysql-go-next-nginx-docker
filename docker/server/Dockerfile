FROM golang:1.16.3-alpine

RUN mkdir -p /go/src/github.com/moneyan/mysql_go_nuxt_nginx_docker

WORKDIR /go/src/github.com/moneyan/mysql_go_nuxt_nginx_docker

ADD ./server /go/src/github.com/moneyan/mysql_go_nuxt_nginx_docker

RUN cd /go/src/github.com/moneyan/mysql_go_nuxt_nginx_docker
RUN go get -d
RUN go get golang.org/x/tools/cmd/godoc