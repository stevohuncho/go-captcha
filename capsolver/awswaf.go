package capsolver

type awsWafPayload struct {
	Type           awsWafType `json:"type"`
	WebsiteURL     string     `json:"websiteURL"`
	AwsKey         string     `json:"awsKey,omitempty"`
	AwsIv          string     `json:"awsIv,omitempty"`
	AwsContext     string     `json:"awsContext,omitempty"`
	AwsChallengeJS string     `json:"awsChallengeJS,omitempty"`
	Proxy          string     `json:"proxy,omitempty"`
}

type awsWafType string

const (
	antiAwsWafTask          awsWafType = "AntiAwsWafTask"
	antiAwsWafTaskProxyLess awsWafType = "AntiAwsWafTaskProxyLess"
)

type awsWafOpts struct {
	awsKey         string
	awsIv          string
	awsContext     string
	awsChallengeJS string
	proxy          *proxy
}

type awsWafOptFunc func(*awsWafOpts)

func defaultAwsWafOpts() *awsWafOpts {
	return &awsWafOpts{
		awsKey:         "",
		awsIv:          "",
		awsContext:     "",
		awsChallengeJS: "",
	}
}

func AwsKeyAwsWafOpt(awsKey string) awsWafOptFunc {
	return func(opts *awsWafOpts) {
		opts.awsKey = awsKey
	}
}

func AwsIvAwsWafOpt(awsIv string) awsWafOptFunc {
	return func(opts *awsWafOpts) {
		opts.awsIv = awsIv
	}
}

func AwsContextAwsWafOpt(awsContext string) awsWafOptFunc {
	return func(opts *awsWafOpts) {
		opts.awsContext = awsContext
	}
}

func AwsChallengeJSAwsWafOpt(awsChallengeJS string) awsWafOptFunc {
	return func(opts *awsWafOpts) {
		opts.awsChallengeJS = awsChallengeJS
	}
}

func ProxyAwsWafOpt(proxy *proxy) awsWafOptFunc {
	return func(opts *awsWafOpts) {
		opts.proxy = proxy
	}
}

func AwsWaf(siteUrl string, opts ...awsWafOptFunc) Payload {
	cfg := defaultAwsWafOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &awsWafPayload{
		WebsiteURL:     siteUrl,
		AwsKey:         cfg.awsKey,
		AwsIv:          cfg.awsIv,
		AwsContext:     cfg.awsContext,
		AwsChallengeJS: cfg.awsChallengeJS,
	}
	if cfg.proxy != nil {
		payload.Type = antiAwsWafTask
		payload.Proxy = cfg.proxy.parse()
	} else {
		payload.Type = antiAwsWafTaskProxyLess
	}
	return payload
}
