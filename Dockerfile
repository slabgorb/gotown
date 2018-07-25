FROM golang:1.8

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/github.com/slabgorb/gotown
COPY . /go/src/github.com/slabgorb/gotown
RUN go get -d -v ./...
RUN go install -v ./...

# CMD if [ ${APP_ENV} = production ]; \
#   then \
#   gotown; \
#   else \
#   go get github.com/codegangsta/gin && \
#   gin --port 3001 --appPort 8003 -i run main.go; \
#   fi

CMD gin --port 3001 --appPort 8003 -i run main.go

EXPOSE 8003 
