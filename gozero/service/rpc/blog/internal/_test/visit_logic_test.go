package _test_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/tracex"
	websiterpclogic "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/logic/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/headerconst"
)

func TestAddVisit(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	md := metadata.MD{}
	md.Set(headerconst.HeaderUid, "1")
	md.Set(headerconst.HeaderTerminal, "terminal")
	md.Set(headerconst.HeaderUserAgent, "")
	ctx = metadata.NewIncomingContext(ctx, md)

	in := &websiterpc.EmptyReq{}

	out, err := websiterpclogic.NewReportLogic(ctx, tsc).Report(in)
	assert.Equal(t, nil, err)
	log.Println(out)
}

func TestGetTotalVisit(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	in := &websiterpc.EmptyReq{}

	out, err := websiterpclogic.NewGetUserTotalVisitLogic(ctx, tsc).GetUserTotalVisit(in)
	assert.Equal(t, nil, err)

	log.Println(out)
}
