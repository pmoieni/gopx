package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	rp *httputil.ReverseProxy
}

func New(origin *url.URL) *Proxy {
	rp := httputil.NewSingleHostReverseProxy(origin)

	rp.Director = func(r *http.Request) {
		r.Header.Add("X-Forwarded-Host", r.Host)
		r.Header.Add("X-Origin-Host", origin.Host)
		r.URL.Scheme = origin.Scheme
		r.URL.Host = origin.Host
		r.URL.Path = origin.Path
	}

	return &Proxy{rp}
}

func (p *Proxy) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.rp.ServeHTTP(w, r)
	}
}
