# News-API
[![Build Status](https://travis-ci.org/IIC2173-2015-2-Grupo2/news-api.svg)](https://travis-ci.org/IIC2173-2015-2-Grupo2/news-api) [![Code Climate](https://codeclimate.com/github/IIC2173-2015-2-Grupo2/news-api/badges/gpa.svg)](https://codeclimate.com/github/IIC2173-2015-2-Grupo2/news-api)

## API Documentation

### Objects

#### `NewsItem`

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

#### `api/v1/private/news`

Returns the new's list

| Argument | Type | Description |
|----------|------|-------------|
|          |      |             |

#### `api/v1/private/news/:id`

Returns the New associated with that `id`

| Argument | Type | Description |
|----------|------|-------------|
|          |      |             |

#### `api/v1/private/news/search`

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

#### Run
Build and run:
```sh
$ make docker
```
