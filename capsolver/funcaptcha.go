package capsolver

type funCaptchaPayload struct {
	Type                     funCaptchaType `json:"type"`
	WebsiteURL               string         `json:"websiteURL"`
	WebsitePublicKey         string         `json:"websitePublicKey"`
	FuncaptchaApiJSSubdomain string         `json:"funcaptchaApiJSSubdomain,omitempty"`
	Data                     string         `json:"data,omitempty"`
}

type funCaptchaType string

const (
	funCaptchaTaskProxyLess funCaptchaType = "FunCaptchaTaskProxyLess"
)

type funCaptchaOpts struct {
	funcaptchaApiJSSubdomain string
	data                     string
}

type funCaptchaOptFunc func(*funCaptchaOpts)

func defaultFunCaptchaOpts() *funCaptchaOpts {
	return &funCaptchaOpts{
		funcaptchaApiJSSubdomain: "",
		data:                     "",
	}
}

func FuncaptchaApiJSSubdomainFunCaptchaOpt(funcaptchaApiJSSubdomain string) funCaptchaOptFunc {
	return func(opts *funCaptchaOpts) {
		opts.funcaptchaApiJSSubdomain = funcaptchaApiJSSubdomain
	}
}

func DataFunCaptchaOpt(data string) funCaptchaOptFunc {
	return func(opts *funCaptchaOpts) {
		opts.data = data
	}
}

func FunCaptcha(siteUrl string, sitePublicKey string, opts ...funCaptchaOptFunc) Payload {
	cfg := defaultFunCaptchaOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &funCaptchaPayload{
		Type:                     funCaptchaTaskProxyLess,
		WebsiteURL:               siteUrl,
		FuncaptchaApiJSSubdomain: cfg.funcaptchaApiJSSubdomain,
		Data:                     cfg.data,
	}
	return payload
}
