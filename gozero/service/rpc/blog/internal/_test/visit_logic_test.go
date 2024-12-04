package _test_test

import (
	"log"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/tracex"
	websiterpclogic "github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/logic/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
)

func TestAddVisit(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	in := &websiterpc.EmptyReq{}

	out, err := websiterpclogic.NewReportLogic(ctx, tsc).Report(in)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out)
}

func TestGetTotalVisit(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	in := &websiterpc.EmptyReq{}

	out, err := websiterpclogic.NewGetUserTotalVisitLogic(ctx, tsc).GetUserTotalVisit(in)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out)
}
