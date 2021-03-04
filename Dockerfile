FROM golang:alpine as build-env
ENV GO111MODULE=on
RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /app
RUN mkdir -p /app/proto
RUN mkdir -p /app/connector
RUN mkdir -p /app/handlers
RUN mkdir -p /app/mockdata
RUN mkdir -p /app/models
RUN mkdir -p /app/config

WORKDIR /app

COPY ./proto/user.pb.go /app/proto
COPY ./connector/connector.go /app/connector
COPY ./handlers/handlers.go /app/handlers
COPY ./mockdata/mockdata.go /app/mockdata
COPY ./models/models.go /app/models
COPY ./server.go /app
COPY ./config/app.yaml /app/config
COPY ./config/config.go /app/config
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app .
CMD ["./app"]

