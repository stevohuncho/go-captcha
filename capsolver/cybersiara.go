package capsolver

type cyberSiAraPayload struct {
	Type             cyberSiAraType `json:"type"`
	WebsiteURL       string         `json:"websiteURL"`
	SlideMasterUrlId string         `json:"SlideMasterUrlId"`
	UserAgent        string         `json:"userAgent"`
	Proxy            string         `json:"proxy,omitempty"`
}

type cyberSiAraType string

const (
	antiCyberSiAraTask          cyberSiAraType = "AntiCyberSiAraTask"
	antiCyberSiAraTaskProxyLess cyberSiAraType = "AntiCyberSiAraTaskProxyLess"
)

type cyberSiAraOpts struct {
	proxy *proxy
}

type cyberSiAraOptFunc func(*cyberSiAraOpts)

func defaultCyberSiAraOpts() *cyberSiAraOpts {
	return &cyberSiAraOpts{
		proxy: nil,
	}
}

func ProxyCyberSiAraOpt(proxy *proxy) cyberSiAraOptFunc {
	return func(opts *cyberSiAraOpts) {
		opts.proxy = proxy
	}
}

func CyberSiAra(siteUrl string, slideMasterUrlId string, userAgent string, opts ...cyberSiAraOptFunc) Payload {
	cfg := defaultCyberSiAraOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &cyberSiAraPayload{
		WebsiteURL:       siteUrl,
		SlideMasterUrlId: slideMasterUrlId,
		UserAgent:        userAgent,
	}
	if cfg.proxy != nil {
		payload.Type = antiCyberSiAraTask
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.Type = antiCyberSiAraTaskProxyLess
	}
	return payload
}
