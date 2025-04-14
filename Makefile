.PHONY: coverage

coverage:
	@echo "Running tests and generating coverage report..."
	go test ./... -coverprofile=coverage.out
	@echo ""
	@echo "Coverage summary per function:"
	go tool cover -func=coverage.out
	@rm -f coverage.out

