package _test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/tracex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/logic/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
)

func Test_getArticleViewCount(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	logic := articlerpclogic.NewArticleHelperLogic(ctx, tsc)

	out := logic.GetArticleViewCount(90)
	log.Println(out)

}

func Test_FindUserLikeArticle(t *testing.T) {
	tsc := svc.NewTestServiceContext()
	ctx := tracex.NewRandomTraceContext()

	in := &articlerpc.UserIdReq{
		UserId: "61ef925a-acd9-4209-bda1-8e313aa279c0",
	}

	out, err := articlerpclogic.NewFindUserLikeArticleLogic(ctx, tsc).FindUserLikeArticle(in)
	assert.Equal(t, nil, err)
	log.Println(out)

}
