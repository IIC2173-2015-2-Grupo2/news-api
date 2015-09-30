default: build

# App name
NAME = news-api
SERVERNAME = $(NAME)-server

# Flags
DBHOST = db-host
DBPORT = db-port

# Binary name
EXEC = ./$(NAME)

get-deps:
	godep go install


###############
# LOCAL
###############

# Remove output
clean:
	rm $(EXEC)

# Build binary
build:
	godep go build

# Start webserver
run:
	$(EXEC) -db-user=$(NEO4USER) -db-password=$(NEO4PASSWORD) -db-host=$(NEO4JHOST) -db-port=$(NEO4JPORT)

test:
	make build
	godep go test -v

# Build and start webserver
start:
	make build
	make run


###############
# DOCKER
###############

# Build docker image
docker-build:
	docker build -t $(NAME) .

# Start application on port 6060
docker-run:
	docker run --publish 6060:8000 --name $(NAME) --rm $(NAME)

# Build and run
docker:
	make docker-build
	make docker-run
