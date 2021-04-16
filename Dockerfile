FROM golang:1.12-alpine

RUN apk add --no-cache git

RUN mkdir /psm-api
ADD . /psm-api
WORKDIR /psm-api

# RUN go mod download

# RUN go mod init br.com.jonathan/psm

# RUN go mod tidy

 # RUN go get github.com/gorilla/mux

# RUN go mod vendor

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["/psm-api/main"]
