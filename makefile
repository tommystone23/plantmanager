FRONTEND_DIR := frontend
BACKEND_DIR := backend
APPLICATION_DIR := applicationservice
BUILD_DIR := build
BINARY_NAME := plantmanager

.PHONY: all
all: frontend backend

# React frontend build
.PHONY: frontend
frontend:
	@echo "Building React frontend..."
	cd $(FRONTEND_DIR) && npm install
	cd $(FRONTEND_DIR) && npm run build
	@echo "Frontend build complete."

# Go backend build
.PHONY: backend
backend:
	@echo "Building Go backend..."
	cd $(BACKEND_DIR)/$(APPLICATION_DIR) && go build -o ../../$(BUILD_DIR)/backend/$(APPLICATION_DIR)/$(BINARY_NAME) main.go
	@echo "Go backend build complete."

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	cd $(FRONTEND_DIR) && rm -rf node_modules
	@echo "Clean complete."

# Run server
.PHONY: run
run: all
	./$(BUILD_DIR)/backend/$(APPLICATION_DIR)/$(BINARY_NAME)
