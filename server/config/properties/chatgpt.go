package properties

type ChatGPT struct {
	ApiHost string `mapstructure:"api_host" json:"api_host" yaml:"api_host"`
	ApiKey  string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`
	Model   string `mapstructure:"model" json:"model" yaml:"model"`
}
