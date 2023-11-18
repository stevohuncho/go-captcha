package capsolver

import "fmt"

type proxy struct {
	url      string
	isSocks5 bool
}

type proxyOptFunc func(*proxy)

func IsSocks5Proxy() proxyOptFunc {
	return func(p *proxy) {
		p.isSocks5 = true
	}
}

func Proxy(url string, opts ...proxyOptFunc) *proxy {
	p := &proxy{
		url:      url,
		isSocks5: false,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *proxy) parse() string {
	if !p.isSocks5 {
		return fmt.Sprintf("http:%s", p.url)
	} else {
		return fmt.Sprintf("socks5:%s", p.url)
	}
}
