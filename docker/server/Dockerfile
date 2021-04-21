FROM golang:1.16.3-stretch
RUN apt update && apt upgrade -y

ENV ROOT=/go/src/server
WORKDIR ${ROOT}

COPY ./server/go.mod ./
COPY ./server/go.sum ./
RUN go mod download

COPY ./server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.8.0/wait /wait
RUN chmod +x /wait

CMD /wait && /go/bin/app
