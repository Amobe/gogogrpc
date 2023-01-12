package abcd

import (
	"testing"

	"github.com/jajajajaja15/gogogrpc/gen/pb/gogogrpc"
	"github.com/stretchr/testify/assert"
)

func TestPayload(t *testing.T) {
	got := fixedPayload()
	want := &gogogrpc.HelloRequest{
		Name: "lex",
	}
	assert.Equal(t, want, got)
}

func BenchmarkPayload(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &gogogrpc.HelloRequest{
			Name: "lex",
		}
	}
}
