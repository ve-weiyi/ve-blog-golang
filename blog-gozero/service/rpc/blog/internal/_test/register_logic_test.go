package _test_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/tracex"
	accountrpclogic "github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/logic/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
)

func TestTxRegister(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()
	in := &accountrpc.RegisterReq{
		Username: "791422175@qq.com",
		Password: "791422171@qq.com",
	}

	register, err := accountrpclogic.NewRegisterLogic(ctx, tsc).Register(in)
	assert.Equal(t, nil, err)

	log.Println(register)
}
