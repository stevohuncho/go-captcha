package capsolver

type cfChallengePayload struct {
	Type       cfChallengeType `json:"type"`
	WebsiteURL string          `json:"websiteURL"`
	Proxy      string          `json:"proxy"`
}

type cfChallengeType string

const (
	antiCloudflareChallengeTask cfChallengeType = "AntiCloudflareTask"
)

func CFChallenge(siteUrl string, proxy *proxy) Payload {
	payload := &cfChallengePayload{
		Type:       antiCloudflareChallengeTask,
		WebsiteURL: siteUrl,
		Proxy:      proxy.parse(),
	}
	return payload
}
