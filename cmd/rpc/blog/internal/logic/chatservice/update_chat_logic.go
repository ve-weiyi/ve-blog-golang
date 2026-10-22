package chatservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

type UpdateChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChatLogic {
	return &UpdateChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新聊天记录
func (l *UpdateChatLogic) UpdateChat(in *chatrpc.UpdateChatRequest) (*chatrpc.UpdateChatResponse, error) {
	uid, _ := metax.GetUserIdFromCtx(l.ctx)
	deviceId, _ := metax.GetDeviceIdFromCtx(l.ctx)
	if uid == "" && deviceId == "" {
		return nil, bizerr.NewBizError(bizcode.CodeUnauthenticated, "缺少用户身份")
	}

	switch in.Status {
	case enums.ChatStatusNormal, enums.ChatStatusEdited, enums.ChatStatusRecall:
	default:
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "聊天状态不合法")
	}

	conditions := "id = ?"
	args := []any{in.Id}
	if uid != "" {
		conditions += " AND user_id = ?"
		args = append(args, uid)
	} else {
		conditions += " AND device_id = ?"
		args = append(args, deviceId)
	}

	rows, err := l.svcCtx.TChatModel.UpdateFields(l.ctx, map[string]interface{}{
		"type":    in.Type,
		"content": in.Content,
		"status":  in.Status,
	}, conditions, args...)
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, bizerr.NewBizError(bizcode.CodeNoPermission, "无权修改该聊天记录")
	}

	return &chatrpc.UpdateChatResponse{Success: true}, nil
}
