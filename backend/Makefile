# Variables
BUF = buf
PROTO_DIR = ./protos
OUT_DIR = ./services

# Generate gRPC code using buf
.PHONY: all clean

all:
	$(BUF) generate $(PROTO_DIR) --output $(OUT_DIR)

clean:
	rm -rf $(OUT_DIR)/*.pb.go
