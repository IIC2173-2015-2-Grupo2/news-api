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
	$(EXEC)

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
	docker build --no-cache --rm --tag=$(NAME) .

# Start application on port 6060
docker-run:
	docker run -e NEO4J_HOST -e NEO4J_PORT -e NEO4J_USER -e NEO4J_PASS --publish 6060:8000 --rm --name $(NAME) $(NAME)

# Build and run
docker:
	make docker-build
	make docker-run
