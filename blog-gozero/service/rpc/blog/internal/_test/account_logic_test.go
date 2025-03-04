package _test_test

import (
	"log"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/tracex"
	accountrpclogic "github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/logic/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
)

func TestGetUserAreasAnalysis(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()
	in := &accountrpc.EmptyReq{}

	out, err := accountrpclogic.NewGetUserAreasAnalysisLogic(ctx, tsc).GetUserAreasAnalysis(in)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out)
}
