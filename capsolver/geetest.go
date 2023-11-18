package capsolver

type geeTestPayload struct {
	Type                      geeTestType `json:"type"`
	WebsiteURL                string      `json:"websiteURL"`
	Gt                        string      `json:"gt"`
	Challenge                 string      `json:"challenge"`
	CaptchaId                 string      `json:"captchaId,omitempty"`
	GeetestApiServerSubdomain string      `json:"geetestApiServerSubdomain,omitempty"`
	Proxy                     string      `json:"proxy,omitempty"`
}

type geeTestType string

const (
	geeTestTask          geeTestType = "GeeTestTask"
	geeTestTaskProxyLess geeTestType = "GeeTestTaskProxyLess"
)

type geeTestOpts struct {
	captchaId                 string
	geetestApiServerSubdomain string
	proxy                     *proxy
}

type geeTestOptFunc func(*geeTestOpts)

func defaultGeeTestOpts() *geeTestOpts {
	return &geeTestOpts{
		captchaId:                 "",
		geetestApiServerSubdomain: "",
		proxy:                     nil,
	}
}

func CaptchaIdGeeTestOpt(captchaId string) geeTestOptFunc {
	return func(opts *geeTestOpts) {
		opts.captchaId = captchaId
	}
}

func GeetestApiServerSubdomainmszGeeTestOpt(geetestApiServerSubdomain string) geeTestOptFunc {
	return func(opts *geeTestOpts) {
		opts.geetestApiServerSubdomain = geetestApiServerSubdomain
	}
}

func ProxyGeeTestOpt(proxy *proxy) geeTestOptFunc {
	return func(opts *geeTestOpts) {
		opts.proxy = proxy
	}
}

func GeeTest(siteUrl string, gt string, challenge string, opts ...geeTestOptFunc) Payload {
	cfg := defaultGeeTestOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &geeTestPayload{
		WebsiteURL:                siteUrl,
		Gt:                        gt,
		Challenge:                 challenge,
		CaptchaId:                 cfg.captchaId,
		GeetestApiServerSubdomain: cfg.geetestApiServerSubdomain,
	}
	if cfg.proxy != nil {
		payload.Type = geeTestTask
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.Type = geeTestTaskProxyLess
	}
	return payload
}
