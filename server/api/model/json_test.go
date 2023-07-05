package model

import (
	"log"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

func TestJson(t *testing.T) {

	//log.Println("\n", jsonconv.ObjectToJsonSnakeIdent(&response.UserMenu{}))

	jstr := `{"id":1,"rolePid":0,"roleDomain":"blog","roleName":"管理员","roleComment":"admin","isDisable":1,"isDefault":false,"createdAt":"2021-03-22T14:10:21+08:00","updatedAt":"2023-05-16T20:31:16+08:00"}`

	var data entity.Role

	err := jsonconv.UnmarshalJSONIgnoreCase([]byte(jstr), &data)
	if err != nil {
		return
	}

	log.Println(jsonconv.ObjectToJsonIndent(data))
}

func TestSnake(t *testing.T) {
	var req map[string]interface{}
	jb := []byte(`{"alipayQRCode":"https://veport.oss-cn-beijing.aliyuncs.com/config/17f234dc487c1bb5bbb732869be0eb53.jpg","gitee":"https://gitee.com/wy791422171","github":"https://github.com/7914-ve","isChatRoom":1,"isCommentReview":0,"isEmailNotice":1,"isMessageReview":0,"isMusicPlayer":0,"isReward":1,"qq":"791422171","socialLoginList":["qq","weibo"],"socialUrlList":["qq","github","gitee"],"touristAvatar":"https://veport.oss-cn-beijing.aliyuncs.com/config/5bfb96809bee5ba80a36811f0bf1d1ea.gif","userAvatar":"https://veport.oss-cn-beijing.aliyuncs.com/config/041a0d1c7fdfb5a610c307e7e44d4f39.jpg","websiteAuthor":"与梦","websiteAvatar":"https://veport.oss-cn-beijing.aliyuncs.com/config/84aa08357bf6e74fc1d4f33552475f91.gif","websiteCreateTime":"2022-01-19","websiteIntro":"分享美好生活。","websiteName":"静闻弦语","websiteNotice":"用户需要查看、发表文章、修改其他信息请登录后台管理系统。网站后台管理系统->https://ve77.cn/admin。     \n网站搭建问题请联系站长QQ791422171。","websiteRecordNo":"桂ICP备2022000185号-1","websocketUrl":"wss://ve77.cn:8088/api/websocket","weiXinQRCode":"https://veport.oss-cn-beijing.aliyuncs.com/config/6bed8a1130b170546341ece729e8819f.jpg"}`)
	err := jsoniter.Unmarshal(jb, &req)
	if err != nil {
		log.Println(err)
	}

	log.Println(jsonconv.ObjectToJsonSnakeIdent(req))
}

func TestIgnore(t *testing.T) {
	var data entity.Role
	var req interface{}
	jb := []byte(`{"id":1,"rolePid":0,"roleDomain":"blog","roleName":"管理员","roleComment":"admin","isDisable":1,"isDefault":false,"createdAt":"2021-03-22T14:10:21+08:00","updatedAt":"2023-05-16T20:31:16+08:00"}`)

	jsoniter.Unmarshal(jb, &req)
	jsonconv.SetCamelCaseJsonTag(&data)
	err := jsoniter.UnmarshalFromString(jsonconv.ObjectToJsonSnake(req), &data)
	if err != nil {
		log.Println(err)
	}

	log.Println(jsonconv.ObjectToJsonSnake(req))
	log.Println(jsoniter.MarshalToString(data))

	//var data int = 42
	//set(&data)
	//Value(&data)
}
func set1(data interface{}) {
	jsonconv.SetCamelCaseJsonTag(data)
}

func set(data interface{}) {
	Value(data)
}

func Value(v interface{}) {
	value := reflect.ValueOf(v).Elem()
	log.Println("--", reflect.TypeOf(v), value.Kind())
}
