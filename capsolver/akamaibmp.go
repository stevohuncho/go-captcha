package capsolver

type akamaiBmpPayload struct {
	Type        akamaiBmpType `json:"type"`
	PackageName string        `json:"packageName"`
	Version     string        `json:"version,omitempty"`
	DeviceId    string        `json:"deviceId,omitempty"`
	DeviceName  string        `json:"deviceName,omitempty"`
	Count       int64         `json:"count,omitempty"`
	Pow         string        `json:"pow,omitempty"`
}

type akamaiBmpType string

const (
	antiAkamaiBmpTask akamaiBmpType = "AntiAkamaiBmpTask"
)

type akamaiBmpOpts struct {
	version    string
	deviceId   string
	deviceName string
	count      int64
	pow        string
	timeout    int
	polling    int
}

type akamaiBmpOptFunc func(*akamaiBmpOpts)

func defaultAkamaiBmpOpts() *akamaiBmpOpts {
	return &akamaiBmpOpts{
		version:    "",
		deviceId:   "",
		deviceName: "",
		count:      -1,
		pow:        "",
	}
}

func VersionAkamaiBmpOpt(version string) akamaiBmpOptFunc {
	return func(opts *akamaiBmpOpts) {
		opts.version = version
	}
}

func DeviceIdAkamaiBmpOpt(deviceId string) akamaiBmpOptFunc {
	return func(opts *akamaiBmpOpts) {
		opts.deviceId = deviceId
	}
}

func DeviceNameAkamaiBmpOpt(deviceName string) akamaiBmpOptFunc {
	return func(opts *akamaiBmpOpts) {
		opts.deviceName = deviceName
	}
}

func CountAkamaiBmpOpt(count int64) akamaiBmpOptFunc {
	return func(opts *akamaiBmpOpts) {
		opts.count = count
	}
}

func PowAkamaiBmpOpt(pow string) akamaiBmpOptFunc {
	return func(opts *akamaiBmpOpts) {
		opts.pow = pow
	}
}

func AkamaiBmp(packageName string, opts ...akamaiBmpOptFunc) Payload {
	cfg := defaultAkamaiBmpOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	payload := &akamaiBmpPayload{
		Type:        antiAkamaiBmpTask,
		PackageName: packageName,
		Version:     cfg.version,
		DeviceId:    cfg.deviceId,
		DeviceName:  cfg.deviceName,
		Pow:         cfg.pow,
	}
	if cfg.count > 0 {
		payload.Count = cfg.count
	}
	return payload
}
