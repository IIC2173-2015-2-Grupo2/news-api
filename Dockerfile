FROM golang:1.5.1-wheezy
RUN go get github.com/tools/godep
EXPOSE 8000
