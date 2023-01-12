GEN_PB_DIR := .

proto:
	protoc --go_out=${GEN_PB_DIR} \
	--go-grpc_out=${GEN_PB_DIR} \
	api/gogogrpc/gogogrpc.proto
