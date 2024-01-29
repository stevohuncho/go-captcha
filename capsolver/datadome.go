package capsolver

import (
	"fmt"
	"net/url"
)

type datadomePayload struct {
	Type       datadomeType `json:"type"`
	CaptchaUrl string       `json:"captchaUrl"`
	Proxy      string       `json:"proxy"`
	UserAgent  string       `json:"userAgent"`
}

type datadomeType string

const (
	dataDomeSliderTask datadomeType = "DataDomeSliderTask"
)

func Datadome(siteUrl string, captchaUrl string, proxy *proxy, userAgent string) Payload {
	referer := url.Values{}
	referer.Add("referer", siteUrl)
	payload := &datadomePayload{
		Type:       dataDomeSliderTask,
		CaptchaUrl: fmt.Sprintf("%s&%s", captchaUrl, referer.Encode()),
		Proxy:      proxy.parse(),
		UserAgent:  userAgent,
	}
	return payload
}
