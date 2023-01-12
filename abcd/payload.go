package abcd

import "github.com/jajajajaja15/gogogrpc/gen/pb/gogogrpc"

func fixedPayload() *gogogrpc.HelloRequest {
	return &gogogrpc.HelloRequest{
		Name: "lex",
	}
}
