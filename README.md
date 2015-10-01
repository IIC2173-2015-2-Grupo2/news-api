# News-API
[![Build Status](https://travis-ci.org/IIC2173-2015-2-Grupo2/news-api.svg)](https://travis-ci.org/IIC2173-2015-2-Grupo2/news-api)

## API Documentation

### Objects

#### `New`

Describes a New.

| Field | Type | Description |
|-------|------|-------------|
| `title` | `string` | - |

### API Usage

#### `api/v1/news/:id`

Returns the New associated with that `id`

| Argument | Type | Description |
|----------|------|-------------|
| `short`  | `boolean` | Request a short version of the new |

#### `api/v1/news/search`

Search news with

| Argument | Type | Description |
|----------|------|-------------|
| `text`   | `string` | Search coincidences on new's body |
| `tags`   | `string[]`| - |


## Development

Install [Golang](https://golang.org/).

Make sure to configure `$GOPATH`. For example:
```sh
$ export GOPATH=$HOME/Repositories/go
$ export PATH=$PATH:$GOPATH/bin
```

Get this repository using `go`:
```sh
$ go get github.com/tools/godep
$ go get github.com/IIC2173-2015-2-Grupo2/news-api

# Project directory
$ cd $GOPATH/src/github.com/IIC2173-2015-2-Grupo2/news-api/
```

### Local

Setup database
```sh
# Example values
export NEO4J_HOST="192.168.99.100"
export NEO4J_PORT="7474"
export NEO4J_USER="neo4j"
export NEO4J_PASS="neo4j"
```

Build and run the project locally using:
```sh
$ make start
```

### [Docker](https://www.docker.com/)

#### Install

##### OSX
Make sure you have installed [Homebrew](http://brew.sh/) and [Homebrew-Cask](http://caskroom.io/).
```sh
# Install Homebrew
$ ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# Install Homebrew-cask
$ brew install caskroom/cask/brew-cask

# Install Docker
$ brew cask install virtualbox
$ brew install docker docker-machine
```

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
