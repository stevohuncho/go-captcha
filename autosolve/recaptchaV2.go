package autosolve

type recaptchaV2Opts struct {
	renderParameters map[string]string
	proxy            string
	userAgent        string
	cookies          string
}

type recaptchaV2OptFunc func(*recaptchaV2Opts)

func DefaultRecaptchaV2Opts() *recaptchaV2Opts {
	return &recaptchaV2Opts{
		renderParameters: nil,
		proxy:            "",
		userAgent:        "",
		cookies:          "",
	}
}

func RenderParametersRecaptchaV2Opt(renderParameters map[string]string) recaptchaV2OptFunc {
	return func(opts *recaptchaV2Opts) {
		opts.renderParameters = renderParameters
	}
}

func ProxyRecaptchaV2Opt(proxy string) recaptchaV2OptFunc {
	return func(opts *recaptchaV2Opts) {
		opts.proxy = proxy
	}
}

func UserAgentRecaptchaV2Opt(userAgent string) recaptchaV2OptFunc {
	return func(opts *recaptchaV2Opts) {
		opts.userAgent = userAgent
	}
}

func CookiesRecaptchaV2Opt(cookies string) recaptchaV2OptFunc {
	return func(opts *recaptchaV2Opts) {
		opts.cookies = cookies
	}
}

func RecaptchaV2(siteUrl string, siteKey string) {}
