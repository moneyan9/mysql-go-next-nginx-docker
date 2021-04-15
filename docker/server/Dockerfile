FROM golang:1.16.3-alpine as dev

RUN apk update && apk add git

ENV ROOT=/go/src/github.com/moneyan/mysql_go_nuxt_nginx_docker
WORKDIR ${ROOT}

COPY ./server/go.mod ./
COPY ./server/go.sum ./
RUN go mod download

EXPOSE 8000
