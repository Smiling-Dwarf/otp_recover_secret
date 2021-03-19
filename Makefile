project = smilingdwarf.com/otp/otp_recovery
.PHONY: help
help: ## help: This help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

build: ## build: build the project
	@echo "\n  Building the Protocol Buffer intermediate file"
	@protoc --go_out=. *.proto
	@echo "\n  Building $(project)"
	@go build $(project)

install: ## install: build and install project to $GOPATH/bin
	@build
	@echo '  Installing $(project) to $$GOPATH/bin'
	@go install $(project)

clean: ## clean: Remove installed binary and intermediate proto
	@rm -fr generated $(GOPATH)/bin/otp_recovery
