# Makefile

# 定义变量
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc
PROTO_DIR = src
OUT_DIR = gen

# 目标：生成 .pb.go 文件
generate:
	@echo "Generating .pb.go files..."
	@protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $(PROTO_DIR)/*.proto
	@echo ".pb.go files have been generated."

# 目标：清理生成的文件
clean:
	@echo "Cleaning generated files..."
	@rm -f $(OUT_DIR)/*.pb.go
	@echo "Cleaned up .pb.go files."

# 默认目标
all: generate
