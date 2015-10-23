FROM golang:onbuild
RUN go get github.com/tools/godep
EXPOSE 8000
