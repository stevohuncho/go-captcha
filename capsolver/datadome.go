package capsolver

type datadomePayload struct {
	Type       datadomeType `json:"type"`
	WebsiteURL string       `json:"websiteURL"`
	CaptchaUrl string       `json:"captchaUrl"`
	Proxy      string       `json:"proxy"`
	UserAgent  string       `json:"userAgent"`
}

type datadomeType string

const (
	dataDomeSliderTask datadomeType = "DataDomeSliderTask"
)

func Datadome(siteUrl string, captchaUrl string, proxy *proxy, userAgent string) Payload {
	payload := &datadomePayload{
		Type:       dataDomeSliderTask,
		WebsiteURL: siteUrl,
		CaptchaUrl: captchaUrl,
		Proxy:      proxy.parse(),
		UserAgent:  userAgent,
	}
	return payload
}
