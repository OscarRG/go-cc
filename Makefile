EXECUTABLE = go-cc

CMD_DIR = ./cmd

BIN_DIR = $(if $(GOBIN),$(GOBIN),$(GOPATH)/bin)

all: build

build:
	@echo "Building $(EXECUTABLE)..."
	@go build -o $(EXECUTABLE) $(CMD_DIR)

install: build
	@echo "Installing $(EXECUTABLE)..."
	@mkdir -p $(BIN_DIR)
	@mv $(EXECUTABLE) $(BIN_DIR)/$(EXECUTABLE)

clean:
	@echo "Cleaning up..."
	@go clean
	@rm -f $(GOPATH)/bin/$(EXECUTABLE)

.PHONY: all build install clean
