package captcha

type CaptchaEventType string

const (
	CaptchaEventSuccess CaptchaEventType = "success"
	CaptchaEventError   CaptchaEventType = "error"
	CaptchaEventInfo    CaptchaEventType = "info"
)

type CaptchaEvent struct {
	Type  CaptchaEventType
	Msg   string
	Token string
}
