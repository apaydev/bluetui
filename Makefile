.PHONY: coverage

run:
	go run ./cmd/bluetui/

run-debug:
	@echo "Running in debug mode..."
	export DEBUG=true && go run ./cmd/bluetui/

coverage:
	@echo "Running tests and generating coverage report..."
	go test ./... -coverprofile=coverage.out
	@echo ""
	@echo "Coverage summary per function:"
	go tool cover -func=coverage.out
	@rm -f coverage.out

