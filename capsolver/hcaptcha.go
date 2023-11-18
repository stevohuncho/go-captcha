package capsolver

type hCaptchaPayload struct {
	Type              hCaptchaType      `json:"type"`
	WebsiteURL        string            `json:"websiteURL"`
	WebsiteKey        string            `json:"websiteKey"`
	IsInvisible       bool              `json:"isInvisible,omitempty"`
	Proxy             string            `json:"proxy,omitempty"`
	EnterprisePayload map[string]string `json:"enterprisePayload,omitempty"`
	GetCaptcha        map[string]string `json:"getCaptcha,omitempty"`
	UserAgent         string            `json:"userAgent,omitempty"`
}

type hCaptchaType string

const (
	hCaptchaTask          hCaptchaType = "HCaptchaTask"
	hCaptchaTaskProxyLess hCaptchaType = "HCaptchaTaskProxyLess"
	hCaptchaTurboTask     hCaptchaType = "HCaptchaTurboTask"
)

type hCaptchaOpts struct {
	isInvisible       bool
	proxy             *proxy
	enterprisePayload map[string]string
	getCaptcha        map[string]string
	userAgent         string
}

type hCaptchaOptFunc func(*hCaptchaOpts)

func defaultHCaptchaOpts() *hCaptchaOpts {
	return &hCaptchaOpts{
		isInvisible:       false,
		proxy:             nil,
		enterprisePayload: nil,
		getCaptcha:        nil,
		userAgent:         "",
	}
}

func InvisibleHCaptchaOpt() hCaptchaOptFunc {
	return func(opts *hCaptchaOpts) {
		opts.isInvisible = true
	}
}
func ProxyHCaptchaOpt(proxy *proxy) hCaptchaOptFunc {
	return func(opts *hCaptchaOpts) {
		opts.proxy = proxy
	}
}

func EnterprisePayloadHCaptchaOpt(enterprisePayload map[string]string) hCaptchaOptFunc {
	return func(opts *hCaptchaOpts) {
		opts.enterprisePayload = enterprisePayload
	}
}

func GetCaptchaPayloadHCaptchaOpt(getCaptcha map[string]string) hCaptchaOptFunc {
	return func(opts *hCaptchaOpts) {
		opts.getCaptcha = getCaptcha
	}
}

func UserAgentHCaptchaOpt(userAgent string) hCaptchaOptFunc {
	return func(opts *hCaptchaOpts) {
		opts.userAgent = userAgent
	}
}

func HCaptcha(siteUrl string, siteKey string, opts ...hCaptchaOptFunc) Payload {
	cfg := defaultHCaptchaOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &hCaptchaPayload{
		WebsiteURL:        siteUrl,
		WebsiteKey:        siteKey,
		IsInvisible:       cfg.isInvisible,
		EnterprisePayload: cfg.enterprisePayload,
		GetCaptcha:        cfg.getCaptcha,
		UserAgent:         cfg.userAgent,
	}
	if cfg.proxy != nil {
		payload.Type = hCaptchaTask
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.Type = hCaptchaTaskProxyLess
	}
	return payload
}

type hCaptchaTurboOpts struct {
	isInvisible       bool
	enterprisePayload map[string]string
	getCaptcha        map[string]string
	userAgent         string
}

type hCaptchaTurboOptFunc func(*hCaptchaTurboOpts)

func defaultHCaptchaTurboOpts() *hCaptchaTurboOpts {
	return &hCaptchaTurboOpts{
		isInvisible:       false,
		enterprisePayload: nil,
		getCaptcha:        nil,
		userAgent:         "",
	}
}

func InvisibleHCaptchaTurboOpt() hCaptchaTurboOptFunc {
	return func(opts *hCaptchaTurboOpts) {
		opts.isInvisible = true
	}
}

func EnterprisePayloadHCaptchaTurboOpt(enterprisePayload map[string]string) hCaptchaTurboOptFunc {
	return func(opts *hCaptchaTurboOpts) {
		opts.enterprisePayload = enterprisePayload
	}
}

func GetCaptchaPayloadHCaptchaTurboOpt(getCaptcha map[string]string) hCaptchaTurboOptFunc {
	return func(opts *hCaptchaTurboOpts) {
		opts.getCaptcha = getCaptcha
	}
}

func UserAgentHCaptchaTurboOpt(userAgent string) hCaptchaTurboOptFunc {
	return func(opts *hCaptchaTurboOpts) {
		opts.userAgent = userAgent
	}
}

func HCaptchaTurbo(siteUrl string, siteKey string, proxy *proxy, opts ...hCaptchaTurboOptFunc) Payload {
	cfg := defaultHCaptchaTurboOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &hCaptchaPayload{
		Type:              hCaptchaTurboTask,
		WebsiteURL:        siteUrl,
		WebsiteKey:        siteKey,
		IsInvisible:       cfg.isInvisible,
		Proxy:             proxy.parse(),
		EnterprisePayload: cfg.enterprisePayload,
		GetCaptcha:        cfg.getCaptcha,
		UserAgent:         cfg.userAgent,
	}
	return payload
}
