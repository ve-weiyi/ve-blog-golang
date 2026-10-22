package metax

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/blog-cloud/infra/http/headerx"
)

// 用户id
func GetApiUserIdFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata error")
	}

	if val, ok := md[headerx.HeaderUid]; ok {
		if len(val) > 0 {
			return val[0], nil
		}
	}

	return "", fmt.Errorf("missing rpc meta '%v'", headerx.HeaderUid)
}

// 游客id (即是用户的终端设备id)
func GetApiDeviceIdFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata error")
	}

	if val, ok := md[headerx.HeaderXDeviceId]; ok {
		if len(val) > 0 {
			return val[0], nil
		}
	}

	return "", fmt.Errorf("missing rpc meta '%v'", headerx.HeaderXDeviceId)
}

func GetApiAppNameFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata error")
	}

	if val, ok := md[headerx.HeaderAppName]; ok {
		if len(val) > 0 {
			return val[0], nil
		}
	}

	return "", fmt.Errorf("missing rpc meta '%v'", headerx.HeaderAppName)
}

func GetApiRemoteAgentFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata error")
	}

	if val, ok := md[HeaderRPCRemoteAgent]; ok {
		if len(val) > 0 {
			return val[0], nil
		}
	}

	return "", fmt.Errorf("missing rpc meta '%v'", HeaderRPCRemoteAgent)
}
