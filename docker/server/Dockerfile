FROM golang:1.16.3-stretch
RUN apt update

ENV ROOT=/go/src/server
WORKDIR ${ROOT}

COPY ./server/go.mod ./
COPY ./server/go.sum ./
RUN go mod download

COPY ./server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

CMD ["/go/bin/app"]