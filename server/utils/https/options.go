package https

type Options struct {
	params  map[string]interface{} // ?a=1&b=2
	headers map[string]string      // map[Content-Type:application/x-www-form-urlencoded]
	body    map[string]interface{} // {"a":1,"b":2}

}

type Option func(options *Options)

func newOptions(opts ...Option) Options {
	// 初始订阅相关配置。默认开启自动ack
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Params(params map[string]interface{}) Option {
	return func(o *Options) {
		o.params = params
	}
}

func Headers(headers map[string]string) Option {
	return func(o *Options) {
		o.headers = headers
	}
}

func Body(body map[string]interface{}) Option {
	return func(o *Options) {
		o.body = body
	}
}
