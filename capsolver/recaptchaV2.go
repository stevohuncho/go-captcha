package capsolver

type reCaptchaV2Payload struct {
	Type              reCaptchaV2Type   `json:"type"`
	WebsiteURL        string            `json:"websiteURL"`
	WebsiteKey        string            `json:"websiteKey"`
	EnterprisePayload map[string]string `json:"enterprisePayload,omitempty"`
	IsInvisible       bool              `json:"isInvisible,omitempty"`
	PageAction        string            `json:"pageAction,omitempty"`
	ApiDomain         string            `json:"apiDomain,omitempty"`
	UserAgent         string            `json:"userAgent,omitempty"`
	Cookies           []cookie          `json:"cookies,omitempty"`
	Proxy             string            `json:"proxy,omitempty"`
	Anchor            string            `json:"anchor,omitempty"`
	Reload            string            `json:"reload,omitempty"`
}

type reCaptchaV2Type string

const (
	reCaptchaV2Task                    reCaptchaV2Type = "ReCaptchaV2Task"
	reCaptchaV2TaskProxyLess           reCaptchaV2Type = "ReCaptchaV2TaskProxyLess"
	reCaptchaV2EnterpriseTask          reCaptchaV2Type = "ReCaptchaV2EnterpriseTask"
	reCaptchaV2EnterpriseTaskProxyLess reCaptchaV2Type = "ReCaptchaV2EnterpriseTaskProxyLess"
)

type reCaptchaV2Opts struct {
	isInvisible bool
	pageAction  string
	apiDomain   string
	userAgent   string
	cookies     []cookie
	proxy       *proxy
	anchor      string
	reload      string
}

type reCaptchaV2OptFunc func(*reCaptchaV2Opts)

func defaultReCaptchaV2Opts() *reCaptchaV2Opts {
	return &reCaptchaV2Opts{
		isInvisible: false,
		pageAction:  "",
		apiDomain:   "",
		userAgent:   "",
		cookies:     nil,
		proxy:       nil,
		anchor:      "",
		reload:      "",
	}
}

func InvisibleReCaptchaV2Opt() reCaptchaV2OptFunc {
	return func(opts *reCaptchaV2Opts) {
		opts.isInvisible = true
	}
}

func PageActionReCaptchaV2Opt(pageAction string) reCaptchaV2OptFunc {
	return func(opts *reCaptchaV2Opts) {
		opts.pageAction = pageAction
	}
}

func ApiDomainReCaptchaV2Opt(apiDomain string) reCaptchaV2OptFunc {
	return func(opts *reCaptchaV2Opts) {
		opts.apiDomain = apiDomain
	}
}

func UserAgentReCaptchaV2Opt(userAgent string) reCaptchaV2OptFunc {
	return func(opts *reCaptchaV2Opts) {
		opts.userAgent = userAgent
	}
}

func CookiesReCaptchaV2Opt(cookies []cookie) reCaptchaV2OptFunc {
	return func(opts *reCaptchaV2Opts) {
		opts.cookies = cookies
	}
}

func ProxyReCaptchaV2Opt(proxy *proxy) reCaptchaV2OptFunc {
	return func(opts *reCaptchaV2Opts) {
		opts.proxy = proxy
	}
}

func AnchorReCaptchaV2Opt(anchor string) reCaptchaV2OptFunc {
	return func(opts *reCaptchaV2Opts) {
		opts.anchor = anchor
	}
}

func ReloadReCaptchaV2Opt(reload string) reCaptchaV2OptFunc {
	return func(opts *reCaptchaV2Opts) {
		opts.reload = reload
	}
}

func ReCaptchaV2(siteUrl string, siteKey string, opts ...reCaptchaV2OptFunc) Payload {
	cfg := defaultReCaptchaV2Opts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &reCaptchaV2Payload{
		WebsiteURL:  siteUrl,
		WebsiteKey:  siteKey,
		IsInvisible: cfg.isInvisible,
		PageAction:  cfg.pageAction,
		ApiDomain:   cfg.apiDomain,
		UserAgent:   cfg.userAgent,
		Cookies:     cfg.cookies,
		Anchor:      cfg.anchor,
		Reload:      cfg.reload,
	}
	if cfg.proxy != nil {
		payload.Type = reCaptchaV2Task
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.Type = reCaptchaV2TaskProxyLess
	}
	return payload
}

type reCaptchaV2EnterpriseOpts struct {
	enterprisePayload map[string]string
	isInvisible       bool
	pageAction        string
	apiDomain         string
	userAgent         string
	cookies           []cookie
	proxy             *proxy
	anchor            string
	reload            string
}

type reCaptchaV2EnterpriseOptFunc func(*reCaptchaV2EnterpriseOpts)

func defaultReCaptchaV2EnterpriseOpts() *reCaptchaV2EnterpriseOpts {
	return &reCaptchaV2EnterpriseOpts{
		enterprisePayload: nil,
		isInvisible:       false,
		pageAction:        "",
		apiDomain:         "",
		userAgent:         "",
		cookies:           nil,
		proxy:             nil,
		anchor:            "",
		reload:            "",
	}
}

func EnterprisePayloadReCaptchaV2EnterpriseOpt(enterprisePayload map[string]string) reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.enterprisePayload = enterprisePayload
	}
}

func InvisibleReCaptchaV2EnterpriseOpt() reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.isInvisible = true
	}
}

func PageActionReCaptchaV2EnterpriseOpt(pageAction string) reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.pageAction = pageAction
	}
}

func ApiDomainReCaptchaV2EnterpriseOpt(apiDomain string) reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.apiDomain = apiDomain
	}
}

func UserAgentReCaptchaV2EnterpriseOpt(userAgent string) reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.userAgent = userAgent
	}
}

func CookiesReCaptchaV2EnterpriseOpt(cookies []cookie) reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.cookies = cookies
	}
}

func ProxyReCaptchaV2EnterpriseOpt(proxy *proxy) reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.proxy = proxy
	}
}

func AnchorReCaptchaV2EnterpriseOpt(anchor string) reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.anchor = anchor
	}
}

func ReloadReCaptchaV2EnterpriseOpt(reload string) reCaptchaV2EnterpriseOptFunc {
	return func(opts *reCaptchaV2EnterpriseOpts) {
		opts.reload = reload
	}
}

func ReCaptchaV2Enterprise(siteUrl string, siteKey string, opts ...reCaptchaV2EnterpriseOptFunc) Payload {
	cfg := defaultReCaptchaV2EnterpriseOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &reCaptchaV2Payload{
		WebsiteURL:        siteUrl,
		WebsiteKey:        siteKey,
		EnterprisePayload: cfg.enterprisePayload,
		IsInvisible:       cfg.isInvisible,
		PageAction:        cfg.pageAction,
		ApiDomain:         cfg.apiDomain,
		UserAgent:         cfg.userAgent,
		Cookies:           cfg.cookies,
		Anchor:            cfg.anchor,
		Reload:            cfg.reload,
	}
	if cfg.proxy != nil {
		payload.Type = reCaptchaV2EnterpriseTask
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.Type = reCaptchaV2EnterpriseTaskProxyLess
	}
	return payload
}
