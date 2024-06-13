package _test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
)

const dsn = "root:mysql7914@(127.0.0.1:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB

func init() {
	var err error
	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "tb_",
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	log.Println("mysql connection done")
}

func Test_Api_Update(t *testing.T) {

	OperationLogModel := model.NewOperationLogModel(db, nil)
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

	batch, err := OperationLogModel.DeleteBatch(ctx, "1=1")
	assert.Equal(t, nil, err)
	t.Log(batch)

	insert, err := OperationLogModel.Insert(ctx, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), insert)
	t.Log(jsonconv.ObjectToJsonIndent(data))

	data.Nickname = "test_nickname_update"
	data.IpAddress = ""
	data.ResponseStatus = 0
	update, err := OperationLogModel.Update(ctx, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), update)

	one, err := OperationLogModel.FindOne(ctx, data.Id)
	assert.Equal(t, nil, err)
	t.Log(jsonconv.ObjectToJsonIndent(one))

	data.Nickname = "test_nickname_save"
	data.IpAddress = ""
	data.ResponseStatus = 0
	save, err := OperationLogModel.Save(ctx, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), save)

	first, err := OperationLogModel.First(ctx, "id = ?", data.Id)
	assert.Equal(t, nil, err)
	t.Log(jsonconv.ObjectToJsonIndent(first))
}
