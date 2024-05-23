package oauth

type Option func(*AuthConfig)

func WithClientID(clientID string) Option {
	return func(o *AuthConfig) {
		o.ClientID = clientID
	}
}

func WithClientSecret(clientSecret string) Option {
	return func(o *AuthConfig) {
		o.ClientSecret = clientSecret
	}
}

func WithRedirectUri(redirectUri string) Option {
	return func(o *AuthConfig) {
		o.RedirectUri = redirectUri
	}
}
