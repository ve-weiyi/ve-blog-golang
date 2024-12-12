package _test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/tracex"
	permissionrpclogic "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/logic/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
)

func Test_FindApiList(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	md := metadata.MD{}
	md.Set(restx.HeaderUid, "1")
	md.Set(restx.HeaderTerminal, "terminal")
	md.Set(restx.HeaderUserAgent, "")
	ctx = metadata.NewIncomingContext(ctx, md)

	in := &permissionrpc.FindApiListReq{}

	out, err := permissionrpclogic.NewFindApiListLogic(ctx, tsc).FindApiList(in)
	assert.Equal(t, nil, err)
	t.Log(jsonconv.AnyToJsonIndent(out))
}
