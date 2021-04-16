FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.16

ENV ROOT=/go/src/github.com/moneyan/mysql_go_nuxt_nginx_docker/server
WORKDIR ${ROOT}

COPY ./server/go.mod ./
COPY ./server/go.sum ./
RUN go mod download

EXPOSE 8000
