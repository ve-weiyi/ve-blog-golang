package websocket

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/ws"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

// WebSocket消息
func WebSocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 接收消息
		receive := func(msg []byte) (tx []byte, err error) {
			logx.Info(string(msg))

			var cs types.ChatSocketMsg
			err = jsonconv.JsonToObject(string(msg), &cs)
			if err != nil {
				return nil, err
			}

			if cs.Content == "" {
				return nil, fmt.Errorf("content is empty")
			}

			uid := cast.ToInt64(r.Context().Value("userId"))
			info, err := svcCtx.UserRpc.FindUserInfo(r.Context(), &blog.UserReq{UserId: uid})
			if err != nil {
				return nil, err
			}

			host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
			if err != nil {
				return nil, err
			}

			ip, err := ipx.GetIpInfoByBaidu(host)
			if err != nil {
				return nil, err
			}

			chat := &blog.ChatRecord{
				UserId:    uid,
				Nickname:  info.Nickname,
				Avatar:    info.Avatar,
				IpAddress: ip.Origip,
				IpSource:  ip.Location,
				Content:   cs.Content,
				Type:      cs.Type,
				CreatedAt: time.Now().Unix(),
			}

			out, err := svcCtx.ChatRpc.CreateChatRecord(r.Context(), chat)
			if err != nil {
				return nil, err
			}

			return []byte(jsonconv.ObjectToJson(out)), nil
		}

		ws.HandleWebSocket(w, r, receive)
	}
}
