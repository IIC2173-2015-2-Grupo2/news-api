# News-API

## Getting Started

Install [Golang](https://golang.org/).

Make sure to configure `$GOPATH`. For example:
```sh
$ export GOPATH=$HOME/Repositories/go
$ export PATH=$PATH:$GOPATH/bin
```

Get this repository using `go`:
```sh
$ go get github.com/IIC2173-2015-2-Grupo2/news-api

# Project directory
$ cd $GOPATH/src/github.com/IIC2173-2015-2-Grupo2/news-api/
```

## Development

### Local

Run the project locally using:
```sh
$ make
```

### Docker

#### Create Virtual Machine
```sh
# Create VM
$ docker-machine create --driver virtualbox news-api-server

# Setup
$ eval "$(docker-machine env news-api-server)"
```

#### Run
Run on port `6060`, to see the Virtual Machine IP:
```sh
$ docker-machine ip news-api-server
```

Build and run:
```sh
$ make docker
```
