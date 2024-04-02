package mail

type Option func(*EmailSender)

func WithHost(host string) Option {
	return func(o *EmailSender) {
		o.Host = host
	}
}

func WithPort(port int) Option {
	return func(o *EmailSender) {
		o.Port = port
	}
}

func WithUsername(username string) Option {
	return func(o *EmailSender) {
		o.Username = username
	}
}

func WithPassword(password string) Option {
	return func(o *EmailSender) {
		o.Password = password
	}
}

func WithNickname(nickname string) Option {
	return func(o *EmailSender) {
		o.Nickname = nickname
	}
}

func WithDeliver(deliver []string) Option {
	return func(o *EmailSender) {
		o.Deliver = deliver
	}
}

func WithIsSSL(isSSL bool) Option {
	return func(o *EmailSender) {
		o.IsSSL = isSSL
	}
}
