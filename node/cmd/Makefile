CURRENT_DIR=.
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_DIR=$(CURRENT_DIR)/bin
BINARY_NAME=$(BINARY_DIR)/givo
API_DIR=$(CURRENT_DIR)/../api

all: test build

build: deps $(BUILD_DIR)
	$(GOBUILD)  -o $(BINARY_NAME) -v .

test: 
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	$(BINARY_NAME)

doc: build
	$(BINARY_NAME) -routes > $(API_DIR)/routes.json

deps:
	$(GOMOD) tidy -v

$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)

.PHONY: deps clean run test build
