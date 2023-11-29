package nacos

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type NacosConfig struct {
	IP          string
	Port        uint64 //8848
	UserName    string
	Password    string
	NameSpaceID string
	Group       string
	DataID      string
	LogLevel    string //debug
	Timeout     int64  //ms
}

type NacosReader struct {
	cfg *NacosConfig
}

func New(cfg *NacosConfig) *NacosReader {
	return &NacosReader{cfg: cfg}
}

func (n *NacosReader) Init(listener func(content string) error) error {
	var dataId, group string
	dataId = n.cfg.DataID
	group = n.cfg.Group

	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(n.cfg.IP, n.cfg.Port),
	}

	//create ClientConfig
	cc := constant.NewClientConfig(
		constant.WithUsername(n.cfg.UserName),
		constant.WithPassword(n.cfg.Password),
		constant.WithNamespaceId(n.cfg.NameSpaceID),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("./runtime/logs"),
		constant.WithCacheDir("./cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		Type:   "yaml",
	})
	fmt.Println("GetConfig,config :"+content, err)

	if err != nil {
		return err
	}

	if err = listener(content); err != nil {
		return err
	}
	go func() {
		//Listen config change,key=dataId+group+namespaceId.
		err = client.ListenConfig(vo.ConfigParam{
			DataId: dataId,
			Group:  group,
			OnChange: func(namespace, group, dataId, data string) {
				fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
				if err = listener(data); err != nil {
					fmt.Println("config changed reload failed")
				}
			},
		})
	}()
	return nil
}
