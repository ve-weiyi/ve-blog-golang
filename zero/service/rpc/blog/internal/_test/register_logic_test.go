package _test_test

import (
	"log"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/tracex"
	authrpclogic "github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/authrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

func TestTxRegister(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()
	in := &blog.RegisterReq{
		Username: "791422175@qq.com",
		Password: "791422171@qq.com",
	}

	register, err := authrpclogic.NewRegisterLogic(ctx, tsc).Register(in)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(register)
}
