FROM golang:latest AS development

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . /usr/src/app
WORKDIR /usr/src/app

CMD CompileDaemon -build="go build -o bin/main cmd/main.go" -command="./bin/main" -exclude-dir=".git"

