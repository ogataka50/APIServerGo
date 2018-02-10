.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
generate: ## generate code from proto
	protoc --go_out=plugins=grpc:./pb proto/*

.PHONY: build_server
build_server: ## build server
	go build -o bin/server server.go

.PHONY: build_client
build_client: ## build client
	go build -o bin/client client.go

.PHONY: build
build: ## build server,client
	$(MAKE) build_server
	$(MAKE) build_client
