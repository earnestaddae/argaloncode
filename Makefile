
BINARY_NAME=alganon
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