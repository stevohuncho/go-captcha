package captcha

import (
	"github.com/stevohuncho/go-captcha/autosolve"
	"github.com/stevohuncho/go-captcha/capsolver"
)

type Harvester struct {
	eventStream *chan CaptchaEvent
	*HarvesterOpts
}

type HarvesterCtx struct{}

type HarvesterOpts struct {
	capsolverClient *capsolver.Client
	autosolveClient *autosolve.Client
}

type HarvesterOptFunc func(*HarvesterOpts)

func DefaultHarvesterOpts() *HarvesterOpts {
	return &HarvesterOpts{
		capsolverClient: nil,
	}
}

func CapsolverHarvesterOpt(c *capsolver.Client) HarvesterOptFunc {
	return func(opts *HarvesterOpts) {
		opts.capsolverClient = c
	}
}

func AutosolveHarvesterOpt(c *autosolve.Client) HarvesterOptFunc {
	return func(opts *HarvesterOpts) {
		opts.autosolveClient = c
	}
}

func CreateHarvester(opts ...HarvesterOptFunc) *Harvester {
	evstr := make(chan CaptchaEvent)
	h := &Harvester{
		eventStream:   &evstr,
		HarvesterOpts: DefaultHarvesterOpts(),
	}
	for _, opt := range opts {
		opt(h.HarvesterOpts)
	}
	return h
}

func (h Harvester) EventSteam() *chan CaptchaEvent {
	return h.eventStream
}

func (h Harvester) Info() *HarvesterCtx {
	return nil
}
