package capsolver

type akamaiWebPayload struct {
	Type      akamaiWebType `json:"type"`
	Url       string        `json:"url"`
	Abck      string        `json:"abck,omitempty"`
	Bmsz      string        `json:"bmsz,omitempty"`
	UserAgent string        `json:"userAgent,omitempty"`
}

type akamaiWebType string

const (
	antiAkamaiWebTask akamaiWebType = "AntiAkamaiWebTask"
)

type akamaiWebOpts struct {
	abck      string
	bmsz      string
	userAgent string
}

type akamaiWebOptFunc func(*akamaiWebOpts)

func defaultAkamaiWebOpts() *akamaiWebOpts {
	return &akamaiWebOpts{
		abck:      "",
		bmsz:      "",
		userAgent: "",
	}
}

func AbckAkamaiWebOpt(abck string) akamaiWebOptFunc {
	return func(opts *akamaiWebOpts) {
		opts.abck = abck
	}
}

func BmszAkamaiWebOpt(bmsz string) akamaiWebOptFunc {
	return func(opts *akamaiWebOpts) {
		opts.bmsz = bmsz
	}
}

func UserAgentAkamaiWebOpt(userAgent string) akamaiWebOptFunc {
	return func(opts *akamaiWebOpts) {
		opts.userAgent = userAgent
	}
}

func AkamaiWeb(siteUrl string, opts ...akamaiWebOptFunc) Payload {
	cfg := defaultAkamaiWebOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &akamaiWebPayload{
		Type:      antiAkamaiWebTask,
		Url:       siteUrl,
		Abck:      cfg.abck,
		Bmsz:      cfg.bmsz,
		UserAgent: cfg.userAgent,
	}
	return payload
}
