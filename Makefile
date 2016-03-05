default: test

test:
	go test ./...

lint:
	go vet ./...
	@echo ""
	golint ./...

cyclo:
	-gocyclo -top 10 -avg .

report:
	@echo "misspell"
	@find . -name \*.go | xargs misspell
	@echo ""
	-gocyclo -over 14 -avg .
	@echo ""
	go vet ./...
	@echo ""
	golint ./...

.PHONY: test lint cyclo report
