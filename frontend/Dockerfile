FROM golang:1.23.1-bookworm

WORKDIR /app
COPY . ./


RUN go mod init your-module-name
RUN go get github.com/gin-gonic/gin
RUN go get github.com/glebarez/go-sqlite
# RUN go get http
# RUN go get os
RUN go build main.go

EXPOSE 8080

ENTRYPOINT [ "./main" ]
