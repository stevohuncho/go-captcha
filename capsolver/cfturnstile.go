package capsolver

type cfTurnstilePayload struct {
	Type       cfTurnstileType `json:"type"`
	WebsiteURL string          `json:"websiteURL"`
	WebsiteKey string          `json:"websiteKey"`
	Proxy      string          `json:"proxy"`
}

type cfTurnstileType string

const (
	antiCloudflareTask cfTurnstileType = "AntiCloudflareTask"
)

func CFTurnstile(siteUrl string, siteKey string, proxy *proxy) Payload {
	payload := &cfTurnstilePayload{
		Type:       antiCloudflareTask,
		WebsiteURL: siteUrl,
		WebsiteKey: siteKey,
		Proxy:      proxy.parse(),
	}
	return payload
}
