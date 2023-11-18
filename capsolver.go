package captcha

import (
	"context"
	"fmt"
	"time"

	"github.com/stevohuncho/go-captcha/capsolver"
)

type capsolverSolveOpts struct {
	timeout int
	polling int
}

type capsolverSolveOptsFunc func(*capsolverSolveOpts)

func defaultCapsolverSolveOpts() *capsolverSolveOpts {
	return &capsolverSolveOpts{
		timeout: -1,
		polling: -1,
	}
}

func PollingCapsolverSolveOpt(polling int) capsolverSolveOptsFunc {
	return func(opts *capsolverSolveOpts) {
		opts.polling = polling
	}
}

func TimeoutCapsolverSolveOpt(timeout int) capsolverSolveOptsFunc {
	return func(opts *capsolverSolveOpts) {
		opts.timeout = timeout
	}
}

func (h Harvester) Capsolver(payload capsolver.Payload, opts ...capsolverSolveOptsFunc) (string, error) {
	if h.capsolverClient == nil {
		return "", fmt.Errorf("invalid capsolver client")
	}
	cfg := defaultCapsolverSolveOpts()
	for _, opt := range opts {
		opt(cfg)
	}
	polling := h.capsolverClient.DefaultPolling()
	if cfg.polling > 0 {
		polling = cfg.polling
	}
	timeout := h.capsolverClient.DefaultTimeout()
	if cfg.timeout > 0 {
		timeout = cfg.timeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))
	defer cancel()
	return h.capsolverClient.Solve(payload, polling, ctx)
}

func (h Harvester) CapsolverBalance() (float64, error) {
	if h.capsolverClient == nil {
		return 0, fmt.Errorf("invalid capsolver client")
	}
	return h.capsolverClient.Balance()
}
