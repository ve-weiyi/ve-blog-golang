package oss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Local(t *testing.T) {
	uploader := NewLocal(&Config{
		BucketUrl: "http://localhost:9999",
		BasePath:  "runtime/uploads",
	})

	//// 上传文件
	//url, err := uploader.UploadLocalFile("test.txt", "test", "test.txt")
	//
	//// 验证结果是否符合预期
	//assert.NoError(t, err) // 验证错误是否为 nil
	//t.Log(url)

	files, err := uploader.ListFiles("", 10)
	assert.NoError(t, err) // 验证错误是否为 nil
	t.Log(files)
}

func Test_Qiniu(t *testing.T) {
	uploader := NewQiniu(&Config{
		Zone:            "ZoneHuanan",
		Endpoint:        "s3.cn-south-1.qiniucs.com",
		AccessKeyId:     "",
		AccessKeySecret: "",
		BucketName:      "veweiyi",
		BucketUrl:       "https://static.veweiyi.cn",
		BasePath:        "blog",
	})

	//// 上传文件
	//url, err := uploader.UploadLocalFile("test.txt", "test", "test.txt")
	//
	//// 验证结果是否符合预期
	//assert.NoError(t, err) // 验证错误是否为 nil
	//t.Log(url)

	files, err := uploader.ListFiles("", 10)
	assert.NoError(t, err) // 验证错误是否为 nil
	t.Log(files)
}

func Test_Aliyun(t *testing.T) {
	uploader := NewAliyunOSS(&Config{
		Zone:            "",
		Endpoint:        "oss-cn-shenzhen.aliyuncs.com",
		AccessKeyId:     "",
		AccessKeySecret: "",
		BucketName:      "ve-blog",
		BucketUrl:       "http://ve-blog.oss-cn-shenzhen.aliyuncs.com",
		BasePath:        "blog",
	})

	//// 上传文件
	//url, err := uploader.UploadLocalFile("test.txt", "test", "test.txt")
	//
	//// 验证结果是否符合预期
	//assert.NoError(t, err) // 验证错误是否为 nil
	//t.Log(url)

	files, err := uploader.ListFiles("", 10)
	assert.NoError(t, err) // 验证错误是否为 nil
	t.Log(files)
}
