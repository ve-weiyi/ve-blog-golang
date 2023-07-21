package iputil

import (
	"fmt"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type BaiduIpResp struct {
	Status       string     `json:"status"`
	T            string     `json:"t"`
	SetCacheTime string     `json:"set_cache_time"`
	Data         []*Address `json:"data"`
}

type Address struct {
	ExtendedLocation string `json:"ExtendedLocation"` // 扩展位置信息
	OriginQuery      string `json:"OriginQuery"`      // 查询的原始地址
	Appinfo          string `json:"appinfo"`          // 应用信息
	DispType         int    `json:"disp_type"`        // 显示类型
	Fetchkey         string `json:"fetchkey"`         // 获取键
	Location         string `json:"location"`         // 地址
	Origip           string `json:"origip"`           // 原始 IP 地址
	Origipquery      string `json:"origipquery"`      // 原始 IP 地址查询
	Resourceid       string `json:"resourceid"`       // 资源 ID
	RoleId           int    `json:"role_id"`          // 角色 ID
	ShareImage       int    `json:"shareImage"`       // 分享图片
	ShowLikeShare    int    `json:"showLikeShare"`    // 显示喜欢分享
	Showlamp         string `json:"showlamp"`         // 是否显示灯泡
	Titlecont        string `json:"titlecont"`        // 标题内容
	Tplt             string `json:"tplt"`             // 模板信息
}

// GetIpSource 获取ip对应的城市地区
func GetIpInfoByBaidu(ip string) (*Address, error) {
	resp, err := http.Get(fmt.Sprintf("http://opendata.baidu.com/api.php?query=" + ip + "&co=&resource_id=6006&oe=utf8"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result BaiduIpResp
	if err := jsoniter.Unmarshal(out, &result); err != nil {
		return nil, err
	}

	if len(result.Data) > 0 {
		return result.Data[0], nil
	} else {
		return nil, fmt.Errorf("no data:%v", result)
	}
}
