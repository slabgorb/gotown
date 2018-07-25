FROM golang:alpine AS build-go

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/github.com/slabgorb/gotown
COPY . /go/src/github.com/slabgorb/gotown
RUN go get -d -v ./...
RUN go install -v ./...

CMD gin --port 3001 --appPort 8003 -i run main.go

EXPOSE 8003 
