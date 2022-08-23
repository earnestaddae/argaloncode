
BINARY_NAME=argalon
ENV=development

# ==================================================================================== # 
# HELPERS
# ==================================================================================== #
## help : print message for each command usage
.PHONY: help
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================================================================== # 
# BUILD
# ==================================================================================== #

## build: Build binary
.PHONY: build 
build:
	@echo "Building back end..."
	go build -o=./bin/api/${BINARY_NAME} ./cmd/api
	@echo "Binary built"

## run: builds and runs the application
.PHONY: run
run: build
	@echo "Starting back end..."
	@env ENV=${ENV} ./bin/api/${BINARY_NAME}
	@echo "Back end started!"

## clean: runs go clean and delete binary
.PHONY: clean
clean:
	@echo "Cleaning"
	@go clean 
	@rm ./bin/api/${BINARY_NAME}
	@echo "Clean"

## start: an alian to run 
.PHONY: start 
start: run 

## stop: stops the running application
.PHONY: stop 
stop:
	@echo "Stopping back end..."
	@-pkill -SIGTERM -f "./bin/api/${BINARY_NAME}"
	@echo "Stopped the back end!"

## restart: stops and starts the running application
.PHONY: restart 
restart: stop start

# ==================================================================================== # 
# QUALITY CONTROL
# ==================================================================================== #
## audit: format, vet and test all code
.PHONY: audit
audit:
	@echo "Formating code..."
	go fmt ./...
	@echo "Vetting code..."
	go vet ./...
	staticcheck ./...
	@echo "Running tests..."
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo "Tidying and verifying module dependencies..."
	go mod tidy 
	go mod verify
	@echo "Vendoring dependencies..."
	go mod vendor