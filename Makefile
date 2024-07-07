#!make

# Define the source directories and the output directory
SRC_DIR := ./cmd
BIN_DIR := ./bin

# List of all commands
COMMANDS := bookbuyer bookstore bookwarehouse bookthief tcp-client tcp-echo-server

# Ensure the output directory exists
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Define a rule to build each command
$(COMMANDS): $(BIN_DIR)
	go build -v -o $(BIN_DIR)/$@ $(SRC_DIR)/$@

# Clean up the build output
clean:
	rm -rf $(BIN_DIR)

# Set the default target to build all commands
.PHONY: all $(COMMANDS) clean
all: $(COMMANDS)

# Define the binaries and corresponding Docker image names
BINARIES := bookbuyer bookstore bookwarehouse bookthief
DOCKER_USER := draychev

# Build Docker images for each binary
.PHONY: build-images
build-images: $(BINARIES:%=build-%)

$(BINARIES:%=build-%): build-%:
	@echo "Building Docker image for $*"
	docker build -t $(DOCKER_USER)/$*:latest -f ./dockerfiles/Dockerfile.$* .

# Push Docker images for each binary
.PHONY: push-images
push-images: $(BINARIES:%=push-%)

$(BINARIES:%=push-%): push-%: build-%
	@echo "Pushing Docker image for $* to Docker Hub"
	docker push $(DOCKER_USER)/$*:latest

# Push Docker images to Docker Hub
.PHONY: push-images
push-images: $(BINARIES:%=push-%)

$(BINARIES:%=push-%): push-%: build-%
	@echo "Pushing Docker image for $* to Docker Hub"
	docker push $(DOCKER_USER)/$*:latest

# Cleanup generated Dockerfiles (optional)
.PHONY: clean
clean:
	@echo "Cleaning up generated Dockerfiles"
	rm -f Dockerfile.bookbuyer Dockerfile.bookstore Dockerfile.bookwarehouse Dockerfile.bookthief
