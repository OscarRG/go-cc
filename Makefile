EXECUTABLE = go-cc

CMD_DIR = ./cmd

all: build

build:
	@echo "Building $(EXECUTABLE)..."
	@go build -o $(EXECUTABLE) $(CMD_DIR)

install: build
	@echo "Installing $(EXECUTABLE)..."
	@mv $(EXECUTABLE) $(GOPATH)/bin/$(EXECUTABLE)

clean:
	@echo "Cleaning up..."
	@go clean
	@rm -f $(GOPATH)/bin/$(EXECUTABLE)

.PHONY: all build install clean
