#!make

# Define the source directories and the output directory
SRC_DIR := ./cmd
BIN_DIR := ./bin

# List of all commands
COMMANDS := bookbuyer bookstore bookwarehouse bookthief bookwatcher tcp-client tcp-echo-server

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
