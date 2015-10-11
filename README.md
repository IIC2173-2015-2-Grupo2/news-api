# News-API
[![Build Status](https://travis-ci.org/IIC2173-2015-2-Grupo2/news-api.svg)](https://travis-ci.org/IIC2173-2015-2-Grupo2/news-api)

## API Documentation

### Objects

#### `New`

Describes a New.

| Field | Type | Description |
|-------|------|-------------|
| `title` | `string` | - |
| `url` | `string` | - |

#### `User`

Describes a User.

| Field | Type | Description |
|-------|------|-------------|
| `name` | `string` | - |
| `username` | `string` | Primary key of `users` |
| `email` | `string` | - |
| `password` | `string` | Only to current registered user |

### API Usage

#### `api/v1/auth/signup`

Creates an account. This returns the account session token

| POST Argument | Type | Description |
|----------|------|-------------|
| `name`  | `string`| - |
| `username`  | `string`| - |
| `email` | `string`| - |
| `password`  | `string`| - |

#### `api/v1/auth/token`

Re-new expired token

| POST Argument | Type | Description |
|----------|------|-------------|
| `username`  | `string`| - |
| `password`  | `string`| - |

#### `api/v1/news`

Returns the new's list

| Argument | Type | Description |
|----------|------|-------------|
|          |      |             |

#### `api/v1/news/:id`

Returns the New associated with that `id`

| Argument | Type | Description |
|----------|------|-------------|
|          |      |             |

#### `api/v1/news/search`

Search news with

| Argument | Type | Description |
|----------|------|-------------|
|          |      |             |


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

Setup other environment variables
```sh
export ENVIRONMENT="PRODUCTION"
export SECRET_HASH="SECRET_HASH"
export AUTH="ENABLE"
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
