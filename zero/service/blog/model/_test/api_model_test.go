package _test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/utils/svctest"
)

func Test_Api_Update(t *testing.T) {
	svcCtx := svctest.NewTestServiceContext()
	ctx := context.Background()

	data := &model.OperationLog{
		UserId:         0,
		Nickname:       "test_nickname",
		IpAddress:      "test_ip_address",
		IpSource:       "test_ip_source",
		OptModule:      "test_opt_module",
		OptDesc:        "test_opt_desc",
		RequestUrl:     "test_request_url",
		RequestMethod:  "test_request_method",
		RequestHeader:  "test_request_header",
		RequestData:    "test_request_data",
		ResponseData:   "test_response_data",
		ResponseStatus: 200,
		Cost:           "",
	}

	batch, err := svcCtx.OperationLogModel.DeleteBatch(ctx, "1=1")
	assert.Equal(t, nil, err)
	t.Log(batch)

	insert, err := svcCtx.OperationLogModel.Insert(ctx, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), insert)
	t.Log(jsonconv.ObjectToJsonIndent(data))

	data.Nickname = "test_nickname_update"
	data.IpAddress = ""
	data.ResponseStatus = 0
	update, err := svcCtx.OperationLogModel.Update(ctx, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), update)

	one, err := svcCtx.OperationLogModel.FindOne(ctx, data.Id)
	assert.Equal(t, nil, err)
	t.Log(jsonconv.ObjectToJsonIndent(one))

	data.Nickname = "test_nickname_save"
	data.IpAddress = ""
	data.ResponseStatus = 0
	save, err := svcCtx.OperationLogModel.Save(ctx, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), save)

	first, err := svcCtx.OperationLogModel.First(ctx, "id = ?", data.Id)
	assert.Equal(t, nil, err)
	t.Log(jsonconv.ObjectToJsonIndent(first))
}
