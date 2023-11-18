package capsolver

type mtCaptchaPayload struct {
	Type       mtCaptchaType `json:"type"`
	WebsiteURL string        `json:"websiteURL"`
	WebsiteKey string        `json:"websiteKey"`
	Proxy      string        `json:"proxy,omitempty"`
}

type mtCaptchaType string

const (
	mtCaptchaTask          mtCaptchaType = "MtCaptchaTask"
	mtCaptchaTaskProxyLess mtCaptchaType = "MtCaptchaTaskProxyLess"
)

type mtCaptchaOpts struct {
	proxy *proxy
}

type mtCaptchaOptFunc func(*mtCaptchaOpts)

func defaultMtCaptchaOpts() *mtCaptchaOpts {
	return &mtCaptchaOpts{
		proxy: nil,
	}
}

func ProxyMtCaptchaOpt(proxy *proxy) mtCaptchaOptFunc {
	return func(opts *mtCaptchaOpts) {
		opts.proxy = proxy
	}
}

func MtCaptcha(siteUrl string, siteKey string, opts ...mtCaptchaOptFunc) Payload {
	cfg := defaultMtCaptchaOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &mtCaptchaPayload{
		WebsiteURL: siteUrl,
		WebsiteKey: siteKey,
	}
	if cfg.proxy != nil {
		payload.Type = mtCaptchaTask
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.Type = mtCaptchaTaskProxyLess
	}
	return payload
}
