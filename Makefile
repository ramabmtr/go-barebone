APP_NAME := go-barebone
APP_DIR := $(shell pwd)
BUILD_DIR := $(APP_DIR)/bin

DOCS_GENERATOR_TOOL := $(BUILD_DIR)/swag
DOCS_GENERATOR_TOOL_SOURCE := github.com/swaggo/swag/cmd/swag
DOCS_GENERATOR_TOOL_TAG := v1.8.12

.PHONY: run

run: build
	@$(BUILD_DIR)/$(APP_NAME)

build: dep generate-docs
	@echo "Building app..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="-s -w" -tags=main -o $(BUILD_DIR)/$(APP_NAME)
	@echo "Build completed"

dep:
	@echo "Downloading dependencies..."
	@go mod tidy
	@echo "Dependencies downloaded"

generate-docs: check-docs-generator-tool
	@echo "Generating docs..."
	@/bin/rm -f ./docs/docs.go ./docs/swagger.json ./docs/swagger.yaml
	@$(DOCS_GENERATOR_TOOL) init --generalInfo ./main.go
	@echo "Docs generated"

check-docs-generator-tool:
ifeq (, $(shell which $(DOCS_GENERATOR_TOOL)))
	@echo "Docs generator tool could not be found"
	@make docs-generator-install
else
	@echo "Docs generator tool found"
endif

docs-generator-install:
	@echo "Installing docs generator..."
	@env GOBIN=$(BUILD_DIR) go install github.com/swaggo/swag/cmd/swag@$(DOCS_GENERATOR_TOOL_TAG)
	@echo "Docs generator installed"
