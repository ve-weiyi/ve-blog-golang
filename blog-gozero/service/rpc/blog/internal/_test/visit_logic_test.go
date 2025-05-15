package _test_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/tracex"
	websiterpclogic "github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/logic/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
)

func TestAddVisit(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	md := metadata.MD{}
	md.Set(restx.HeaderUid, "1")
	md.Set(restx.HeaderTerminalId, "terminal")
	md.Set(restx.HeaderUserAgent, "")
	ctx = metadata.NewIncomingContext(ctx, md)

	in := &websiterpc.AddVisitReq{}

	out, err := websiterpclogic.NewAddVisitLogic(ctx, tsc).AddVisit(in)
	assert.Equal(t, nil, err)
	log.Println(out)
}

func TestAnalysisVisit(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	in := &websiterpc.EmptyReq{}

	out, err := websiterpclogic.NewAnalysisVisitLogic(ctx, tsc).AnalysisVisit(in)
	assert.Equal(t, nil, err)

	log.Println(out)
}
