package properties

type Upload struct {
	Mode   string `mapstructure:"mode"`
	Local  Local  `mapstructure:"local"`
	Aliyun Aliyun `mapstructure:"aliyun"`
	Qiniu  Aliyun `mapstructure:"qiniu"`
}

type Local struct {
	Url  string `mapstructure:"url" json:"url" yaml:"url"`    // 本地项目访问地址
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件访问目录
}

type Aliyun struct {
	Zone            string `mapstructure:"zone" json:"zone" yaml:"zone"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"access-key-secret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" json:"bucket-name" yaml:"bucket-name"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucket-url" yaml:"bucket-url"`
	BasePath        string `mapstructure:"base-path" json:"base-path" yaml:"base-path"`
}
