default: build

# App name
NAME = news-api
SERVERNAME = $(NAME)-server

# Binary name
EXEC = ./$(NAME)

get-deps:
	go get github.com/gin-gonic/gin


###############
# LOCAL
###############

# Remove output
clean:
	rm $(EXEC)

# Build binary
build:
	make clean
	go build

# Start webserver
run:
	$(EXEC)

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
