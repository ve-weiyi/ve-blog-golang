package metax

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/constantx"
)

func GetRPCInnerXUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata error")
	}

	if val, ok := md[constantx.HeaderXUserID]; ok {
		if len(val) > 0 {
			return val[0], nil
		}
	}

	return "", fmt.Errorf("get rpc meta error:%v", constantx.HeaderXUserID)
}
