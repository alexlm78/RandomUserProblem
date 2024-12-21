APP_NAME = ruProb
all: build

build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) SolOne.go SolTwo.go main.go

clean:
	@echo "Cleaning..."
	@rm -rf bin/*

run: build
	@echo "Running..."
	@go run bin/$(APP_NAME)

.phony: build clean run
