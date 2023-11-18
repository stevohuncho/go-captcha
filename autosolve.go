package captcha

import "github.com/stevohuncho/go-captcha/autosolve"

type autosolveSolveOpts struct {
	timeout int
	polling int
}

type autosolveSolveOptsFunc func(*autosolveSolveOpts)

func defaultAutosolveSolveOpts() *autosolveSolveOpts {
	return &autosolveSolveOpts{
		timeout: -1,
		polling: -1,
	}
}

func PollingAutosolveSolveOpt(polling int) autosolveSolveOptsFunc {
	return func(opts *autosolveSolveOpts) {
		opts.polling = polling
	}
}

func TimeoutAutosolveSolveOpt(timeout int) autosolveSolveOptsFunc {
	return func(opts *autosolveSolveOpts) {
		opts.timeout = timeout
	}
}

func Autosolve(paylod autosolve.Payload, opts ...autosolveSolveOpts) (string, error) {
	return "", nil
}
