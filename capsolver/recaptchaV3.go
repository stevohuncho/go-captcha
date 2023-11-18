package capsolver

type reCaptchaV3Payload struct {
	TaskType          reCaptchaV3TaskType `json:"type"`
	WebsiteURL        string              `json:"websiteURL"`
	WebsiteKey        string              `json:"websiteKey"`
	PageAction        string              `json:"pageAction"`
	MinScore          reCaptchaV3MinScore `json:"minScore,omitempty"`
	Proxy             string              `json:"proxy,omitempty"`
	EnterprisePayload map[string]string   `json:"enterprisePayload,omitempty"`
	ApiDomain         string              `json:"apiDomain,omitempty"`
	UserAgent         string              `json:"userAgent,omitempty"`
	Cookies           []cookie            `json:"cookies,omitempty"`
	Anchor            string              `json:"anchor,omitempty"`
	Reload            string              `json:"reload,omitempty"`
}

type reCaptchaV3TaskType string

const (
	reCaptchaV3Task                    reCaptchaV3TaskType = "ReCaptchaV3Task"
	reCaptchaV3TaskProxyLess           reCaptchaV3TaskType = "ReCaptchaV3TaskProxyLess"
	reCaptchaV3EnterpriseTask          reCaptchaV3TaskType = "ReCaptchaV3EnterpriseTask"
	reCaptchaV3EnterpriseTaskProxyLess reCaptchaV3TaskType = "ReCaptchaV3EnterpriseTaskProxyLess"
)

type reCaptchaV3MinScore float64

const (
	ReCaptchaV3Point1 reCaptchaV3MinScore = 0.1
	ReCaptchaV3Point2 reCaptchaV3MinScore = 0.2
	ReCaptchaV3Point3 reCaptchaV3MinScore = 0.3
	ReCaptchaV3Point4 reCaptchaV3MinScore = 0.4
	ReCaptchaV3Point5 reCaptchaV3MinScore = 0.5
	ReCaptchaV3Point6 reCaptchaV3MinScore = 0.6
	ReCaptchaV3Point7 reCaptchaV3MinScore = 0.7
	ReCaptchaV3Point8 reCaptchaV3MinScore = 0.8
	ReCaptchaV3Point9 reCaptchaV3MinScore = 0.9
)

type reCaptchaV3Opts struct {
	minScore  reCaptchaV3MinScore
	proxy     *proxy
	apiDomain string
	userAgent string
	cookies   []cookie
	anchor    string
	reload    string
}

type reCaptchaV3OptFunc func(*reCaptchaV3Opts)

func defaultReCaptchaV3Opts() *reCaptchaV3Opts {
	return &reCaptchaV3Opts{
		minScore:  -1,
		proxy:     nil,
		apiDomain: "",
		userAgent: "",
		cookies:   nil,
		anchor:    "",
		reload:    "",
	}
}

func MinScoreReCaptchaV3Opt(minScore reCaptchaV3MinScore) reCaptchaV3OptFunc {
	return func(opts *reCaptchaV3Opts) {
		opts.minScore = minScore
	}
}

func ProxyReCaptchaV3Opt(proxy *proxy) reCaptchaV3OptFunc {
	return func(opts *reCaptchaV3Opts) {
		opts.proxy = proxy
	}
}

func ApiDomainReCaptchaV3Opt(apiDomain string) reCaptchaV3OptFunc {
	return func(opts *reCaptchaV3Opts) {
		opts.apiDomain = apiDomain
	}
}

func UserAgentReCaptchaV3Opt(userAgent string) reCaptchaV3OptFunc {
	return func(opts *reCaptchaV3Opts) {
		opts.userAgent = userAgent
	}
}

func CookiesReCaptchaV3Opt(cookies []cookie) reCaptchaV3OptFunc {
	return func(opts *reCaptchaV3Opts) {
		opts.cookies = cookies
	}
}

func AnchorReCaptchaV3Opt(anchor string) reCaptchaV3OptFunc {
	return func(opts *reCaptchaV3Opts) {
		opts.anchor = anchor
	}
}

func ReloadReCaptchaV3Opt(reload string) reCaptchaV3OptFunc {
	return func(opts *reCaptchaV3Opts) {
		opts.reload = reload
	}
}

func ReCaptchaV3(siteUrl string, siteKey string, pageAction string, opts ...reCaptchaV3OptFunc) Payload {
	cfg := defaultReCaptchaV3Opts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &reCaptchaV3Payload{
		WebsiteURL: siteUrl,
		WebsiteKey: siteKey,
		PageAction: pageAction,
		MinScore:   cfg.minScore,
		ApiDomain:  cfg.apiDomain,
		UserAgent:  cfg.userAgent,
		Cookies:    cfg.cookies,
		Anchor:     cfg.anchor,
		Reload:     cfg.reload,
	}
	if cfg.proxy != nil {
		payload.TaskType = reCaptchaV3Task
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.TaskType = reCaptchaV3TaskProxyLess
	}
	return payload
}

type reCaptchaV3EnterpriseOpts struct {
	minScore          reCaptchaV3MinScore
	proxy             *proxy
	enterprisePayload map[string]string
	apiDomain         string
	userAgent         string
	cookies           []cookie
	anchor            string
	reload            string
}

type reCaptchaV3EnterpriseOptFunc func(*reCaptchaV3EnterpriseOpts)

func defaultReCaptchaV3EnterpriseOpts() *reCaptchaV3EnterpriseOpts {
	return &reCaptchaV3EnterpriseOpts{
		minScore:          -1,
		proxy:             nil,
		enterprisePayload: nil,
		apiDomain:         "",
		userAgent:         "",
		cookies:           nil,
		anchor:            "",
		reload:            "",
	}
}

func EnterprisePayloadReCaptchaV3EnterpriseOpt(enterprisePayload map[string]string) reCaptchaV3EnterpriseOptFunc {
	return func(opts *reCaptchaV3EnterpriseOpts) {
		opts.enterprisePayload = enterprisePayload
	}
}

func MinScoreReCaptchaV3EnterpriseOpt(minScore reCaptchaV3MinScore) reCaptchaV3EnterpriseOptFunc {
	return func(opts *reCaptchaV3EnterpriseOpts) {
		opts.minScore = minScore
	}
}

func ProxyReCaptchaV3EnterpriseOpt(proxy *proxy) reCaptchaV3EnterpriseOptFunc {
	return func(opts *reCaptchaV3EnterpriseOpts) {
		opts.proxy = proxy
	}
}

func ApiDomainReCaptchaV3EnterpriseOpt(apiDomain string) reCaptchaV3EnterpriseOptFunc {
	return func(opts *reCaptchaV3EnterpriseOpts) {
		opts.apiDomain = apiDomain
	}
}

func UserAgentReCaptchaV3EnterpriseOpt(userAgent string) reCaptchaV3EnterpriseOptFunc {
	return func(opts *reCaptchaV3EnterpriseOpts) {
		opts.userAgent = userAgent
	}
}

func CookiesReCaptchaV3EnterpriseOpt(cookies []cookie) reCaptchaV3EnterpriseOptFunc {
	return func(opts *reCaptchaV3EnterpriseOpts) {
		opts.cookies = cookies
	}
}

func AnchorReCaptchaV3EnterpriseOpt(anchor string) reCaptchaV3EnterpriseOptFunc {
	return func(opts *reCaptchaV3EnterpriseOpts) {
		opts.anchor = anchor
	}
}

func ReloadReCaptchaV3EnterpriseOpt(reload string) reCaptchaV3EnterpriseOptFunc {
	return func(opts *reCaptchaV3EnterpriseOpts) {
		opts.reload = reload
	}
}

func ReCaptchaV3Enterprise(siteUrl string, siteKey string, pageAction string, opts ...reCaptchaV3EnterpriseOptFunc) Payload {
	cfg := defaultReCaptchaV3EnterpriseOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &reCaptchaV3Payload{
		WebsiteURL:        siteUrl,
		WebsiteKey:        siteKey,
		PageAction:        pageAction,
		MinScore:          cfg.minScore,
		EnterprisePayload: cfg.enterprisePayload,
		ApiDomain:         cfg.apiDomain,
		UserAgent:         cfg.userAgent,
		Cookies:           cfg.cookies,
		Anchor:            cfg.anchor,
		Reload:            cfg.reload,
	}
	if cfg.proxy != nil {
		payload.TaskType = reCaptchaV3Task
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.TaskType = reCaptchaV3TaskProxyLess
	}
	return payload
}
