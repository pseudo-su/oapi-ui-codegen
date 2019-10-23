help:
	@echo "This is a helper makefile for oapi-ui-codegen"
	@echo "Targets:"
	@echo "    generate:    regenerate all generated files"
	@echo "    test:        run all tests"

.PHONY: generate
generate:
	go generate ./codegen/...
	go generate ./...

.PHONY: test
test:
	go test -cover ./... -count=1
