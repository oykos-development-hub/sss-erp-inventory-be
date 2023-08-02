BINARY_NAME=inventoryApp

build:
	@echo "Building Inventory APP..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Inventory APP built!"

run: build
	@echo "Starting Inventory APP..."
	@./tmp/${BINARY_NAME}
	@echo "Inventory APP started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Inventory APP..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Inventory APP!"

restart: stop start