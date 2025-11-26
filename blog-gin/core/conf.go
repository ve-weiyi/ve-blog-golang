package core

import (
	"bytes"
	"fmt"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/mapstructure/v2"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/config"
)

// ConfigOption 配置选项（用于初始化时传递参数）
type ConfigOption struct {
	ConfigMode string // 运行方式 file|nacos
	ConfigType string // 配置类型（yaml/json等）

	// 本地配置相关
	LocalPath string // 本地配置文件路径（如config.yaml）

	// Nacos相关配置
	NacosHost       string // Nacos服务地址
	NacosPort       uint64 // Nacos服务端口
	NacosNamespace  string // Nacos命名空间
	NacosDataID     string // 配置DataID
	NacosGroup      string // 配置Group
	NacosRuntimeDir string // Nacos运行时目录
	NacosUsername   string // Nacos用户名
	NacosPassword   string // Nacos密码
}

func InitConfig(option *ConfigOption) (*config.Config, error) {
	switch strings.ToLower(option.ConfigMode) {
	case "file":
		return initLocalConfig(option)
	case "nacos":
		return initNacosConfig(option)
	default:
		return nil, fmt.Errorf("unsupported config mode: %s", option.ConfigMode)
	}
}

// initLocalConfig 初始化本地配置并监听文件变化
func initLocalConfig(option *ConfigOption) (*config.Config, error) {
	fmt.Println("Initializing local config from file:", option.LocalPath)

	// 初始化配置指针（后续更新会直接修改该指针指向的内存）
	cfg := &config.Config{}
	v := viper.New()

	v.SetConfigFile(option.LocalPath)
	v.SetConfigType(option.ConfigType)
	// 修改映射字段为json tag
	hook := func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "json"
	}

	// 首次读取本地文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read local config: %w", err)
	}

	// 首次解析到cfg指针
	if err := v.Unmarshal(cfg, hook); err != nil {
		return nil, fmt.Errorf("failed to unmarshal local config: %w", err)
	}

	// 监听本地文件变化：更新时直接修改cfg指向的内存
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Local config file changed: %s\n", e.Name)
		// 重新解析到原指针（关键：不创建新指针，直接更新数据）
		if err := v.Unmarshal(cfg, hook); err != nil {
			fmt.Printf("failed to update local config: %v\n", err)
			return
		}
		fmt.Println("Local config updated successfully")
	})

	fmt.Println("Local config initialization completed")
	fmt.Printf("%+v\n", cfg)
	return cfg, nil // 返回指针，外部持有后会同步更新
}

// initNacosConfig 初始化Nacos配置并监听变化
func initNacosConfig(option *ConfigOption) (*config.Config, error) {
	fmt.Println("Initializing Nacos config from file:", option.NacosHost)
	// 初始化配置指针（后续更新会直接修改该指针指向的内存）
	cfg := &config.Config{}
	v := viper.New()
	v.SetConfigType(option.ConfigType)
	// 修改映射字段为json tag
	hook := func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "json"
	}

	// 1. 创建Nacos客户端配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: option.NacosHost,
			Port:   option.NacosPort, // 默认端口，可根据实际修改
		},
	}

	clientConfig := &constant.ClientConfig{
		Username:            option.NacosUsername,
		Password:            option.NacosPassword,
		NamespaceId:         option.NacosNamespace, // 命名空间ID（默认为public）
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              path.Join(option.NacosRuntimeDir, "logs"),
		CacheDir:            path.Join(option.NacosRuntimeDir, "cache"),
		LogLevel:            "info",
	}

	// 2. 创建Nacos配置客户端
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Nacos client: %w", err)
	}

	// 3. 首次从Nacos获取配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: option.NacosDataID,
		Group:  option.NacosGroup,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get Nacos config: %w", err)
	}

	// 4. 解析Nacos配置（修复：传指针给Unmarshal）
	if err := v.ReadConfig(strings.NewReader(content)); err != nil {
		return nil, fmt.Errorf("failed to read Nacos content: %w", err)
	}
	if err := v.Unmarshal(cfg, hook); err != nil { // 关键：传指针&cfg，原代码漏了&
		return nil, fmt.Errorf("failed to unmarshal Nacos config: %w", err)
	}

	// 5. 监听Nacos配置变化：更新时直接修改cfg指向的内存
	err = client.ListenConfig(vo.ConfigParam{
		DataId: option.NacosDataID,
		Group:  option.NacosGroup,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Printf("Nacos config changed: dataId=%s, group=%s\n", dataId, group)
			// 重新解析变更后的配置到原指针
			if err := v.ReadConfig(bytes.NewReader([]byte(data))); err != nil {
				fmt.Printf("failed to parse updated Nacos config: %v\n", err)
				return
			}

			if err := v.Unmarshal(cfg, hook); err != nil { // 关键：传指针&cfg
				fmt.Printf("failed to unmarshal updated Nacos config: %v\n", err)
				return
			}
			fmt.Println("Nacos config updated successfully")
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to listen Nacos config changes: %w", err)
	}

	fmt.Println("Nacos config initialization completed")
	return cfg, nil // 返回指针，外部持有后会同步更新
}
