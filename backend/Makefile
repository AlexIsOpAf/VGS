# Define Go installation paths for the protoc plugins
GOBIN := $(shell go env GOPATH)/bin

# Define paths for buf and protoc plugins
BUF := $(GOBIN)/buf
BUF_EXE := $(GOBIN)/buf.exe
PROTOC_GEN_GO := $(GOBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(GOBIN)/protoc-gen-go-grpc

# Detect the operating system
UNAME_S := $(shell uname -s)
ARCH := $(shell uname -m)
BUF_VERSION := v1.24.0

# Define the directories for .proto files and the output
PROTO_DIR := ./protos
OUT_DIR := ./services

# Cross-platform commands
.PHONY: deps_linux deps_mac deps_windows generate_linux generate_mac generate_windows clean uninstall_deps generate

# Install Buf for Linux
install_buf_linux:
	@echo "Installing Buf for Linux..."
	curl -sSL https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-Linux-x86_64 -o $(BUF) && chmod +x $(BUF)

# Install Buf for macOS
install_buf_mac:
	@echo "Installing Buf for macOS..."
	@if [ "$(ARCH)" = "arm64" ]; then \
		curl -sSL https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-Darwin-arm64 -o $(BUF) && chmod +x $(BUF); \
	else \
		curl -sSL https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-Darwin-x86_64 -o $(BUF) && chmod +x $(BUF); \
	fi

# Install Buf for Windows
install_buf_windows:
	@echo "Installing Buf for Windows..."
	curl -sSL https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-Windows-x86_64.exe -o $(BUF_EXE) && chmod +x $(BUF_EXE)

# Install Go plugins
install_protoc_plugins:
	@if [ ! -f "$(PROTOC_GEN_GO)" ]; then \
		echo "Installing protoc-gen-go..."; \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
	else \
		echo "protoc-gen-go already installed"; \
	fi
	@if [ ! -f "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "Installing protoc-gen-go-grpc..."; \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest; \
	else \
		echo "protoc-gen-go-grpc already installed"; \
	fi

# Linux dependencies
deps_linux: install_buf_linux install_protoc_plugins

# macOS dependencies
deps_mac: install_buf_mac install_protoc_plugins

# Windows dependencies
deps_windows: install_buf_windows install_protoc_plugins

# Generate gRPC code for Linux
generate_linux: deps_linux
	@echo "Generating gRPC code for Linux..."
	@PATH=$(GOBIN):$$PATH $(BUF) generate $(PROTO_DIR)

# Generate gRPC code for macOS
generate_mac: deps_mac
	@echo "Generating gRPC code for macOS..."
	@PATH=$(GOBIN):$$PATH $(BUF) generate $(PROTO_DIR)

# Generate gRPC code for Windows
generate_windows: deps_windows
	@echo "Generating gRPC code for Windows..."
	@PATH=$(GOBIN):$$PATH $(BUF_EXE) generate $(PROTO_DIR)

# Clean generated files
clean:
	rm -rf $(OUT_DIR)

# Uninstall dependencies
uninstall_deps:
	@echo "Uninstalling Buf and Go plugins..."
	@if [ -f "$(BUF)" ]; then \
		echo "Removing Buf..."; \
		rm -f $(BUF); \
	elif [ -f "$(BUF_EXE)" ]; then \
		echo "Removing Buf (Windows)..."; \
		rm -f $(BUF_EXE); \
	else \
		echo "Buf is not installed"; \
	fi
	@if [ -f "$(PROTOC_GEN_GO)" ]; then \
		echo "Removing protoc-gen-go..."; \
		rm -f $(PROTOC_GEN_GO); \
	else \
		echo "protoc-gen-go is not installed"; \
	fi
	@if [ -f "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "Removing protoc-gen-go-grpc..."; \
		rm -f $(PROTOC_GEN_GO_GRPC); \
	else \
		echo "protoc-gen-go-grpc is not installed"; \
	fi
	@echo "Dependencies have been uninstalled."

# Automatically detect the platform and run the appropriate commands
generate:
	@if [ "$(UNAME_S)" = "Linux" ]; then \
		echo "Detected Linux..."; \
		$(MAKE) generate_linux; \
	elif [ "$(UNAME_S)" = "Darwin" ]; then \
		echo "Detected macOS..."; \
		$(MAKE) generate_mac; \
	elif [ "$(UNAME_S)" = "MINGW32_NT" ] || [ "$(UNAME_S)" = "MINGW64_NT" ]; then \
		echo "Detected Windows..."; \
		$(MAKE) generate_windows; \
	else \
		echo "Unsupported platform: $(UNAME_S)"; \
		exit 1; \
	fi
