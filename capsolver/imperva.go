package capsolver

type impervaPayload struct {
	Type           impervaType `json:"type"`
	WebsiteUrl     string      `json:"websiteUrl"`
	UserAgent      string      `json:"userAgent"`
	Proxy          string      `json:"proxy"`
	Utmvc          bool        `json:"utmvc,omitempty"`
	Reese84        bool        `json:"reese84,omitempty"`
	ReeseScriptUrl string      `json:"reeseScriptUrl,omitempty"`
	Cookies        []cookie    `json:"cookies,omitempty"`
	ReeseToken     string      `json:"reeseToken,omitempty"`
}

type impervaType string

const (
	antiImpervaTask impervaType = "AntiImpervaTask"
)

type impervaOpts struct {
	utmvc          bool
	reese84        bool
	reeseScriptUrl string
	cookies        []cookie
	reeseToken     string
}

type impervaOptFunc func(*impervaOpts)

func defaultImpervaOpts() *impervaOpts {
	return &impervaOpts{
		utmvc:          false,
		reese84:        false,
		reeseScriptUrl: "",
		cookies:        nil,
		reeseToken:     "",
	}
}

func UtmvcImpervaOpt() impervaOptFunc {
	return func(opts *impervaOpts) {
		opts.utmvc = true
	}
}

func Reese84ImpervaOpt() impervaOptFunc {
	return func(opts *impervaOpts) {
		opts.reese84 = true
	}
}

func ReeseScriptUrlImpervaOpt(reeseScriptUrl string) impervaOptFunc {
	return func(opts *impervaOpts) {
		opts.reeseScriptUrl = reeseScriptUrl
	}
}

func CookiesImpervaOpt(cookies []cookie) impervaOptFunc {
	return func(opts *impervaOpts) {
		opts.cookies = cookies
	}
}

func ReeseTokenImpervaOpt(reeseToken string) impervaOptFunc {
	return func(opts *impervaOpts) {
		opts.reeseToken = reeseToken
	}
}

func Imperva(siteUrl string, userAgent string, proxy *proxy, opts ...impervaOptFunc) Payload {
	cfg := defaultImpervaOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &impervaPayload{
		Type:           antiImpervaTask,
		WebsiteUrl:     siteUrl,
		UserAgent:      userAgent,
		Proxy:          proxy.parse(),
		Utmvc:          cfg.utmvc,
		Reese84:        cfg.reese84,
		ReeseScriptUrl: cfg.reeseScriptUrl,
		Cookies:        cfg.cookies,
		ReeseToken:     cfg.reeseToken,
	}
	return payload
}
