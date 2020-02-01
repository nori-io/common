# NoriCommon Makefile

gen-mocks: ## generate mocks
	@mockgen -source=config/manager.go -destination=mocks/mock_config/config_manager.go
	@mockgen -source=logger/logger.go -destination=mocks/mock_logger/logger.go
	@mockgen -source=meta/meta.go -destination=mocks/mock_meta/meta.go
	@mockgen -source=plugin/plugin.go -destination=mocks/mock_plugin/plugin.go
	@mockgen -source=plugin/registry.go -destination=mocks/mock_plugin/registry.go
	@mockgen -source=storage/storage.go -destination=mocks/mock_storage/storage.go
.PHONY: gen-mocks

test: ## run go tests
	@go test -v -race ./...
.PHONY: test

test-with-coverage: ## run go test with coverage
	@go test -race -covermode atomic -coverprofile profile.out ${TEST_ARGS} ./... ;\
	go tool cover -func=profile.out
.PHONY: test-with-coverage

help: ## run this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

.DEFAULT_GOAL := help