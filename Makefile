.DEFAULT_GOAL := help

.PHONY: test
test: ## Run go test
	go test -cover .

.PHONY: gen
gen: ## Generate sample csv and json
	go run ./cmd/iso3166ja -f csv -o iso3166ja
	go run ./cmd/iso3166ja -f csv -o iso3166ja-2 -c alpha2,name,name_ja
	go run ./cmd/iso3166ja -f json-array -o iso3166ja-a
	go run ./cmd/iso3166ja -f json-map -o iso3166ja-m

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
