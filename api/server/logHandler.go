package server

import (
	"lca/internal/pkg/logging"
	"lca/internal/pkg/util"
	"net/http"
	"strings"
	"time"

	"github.com/felixge/httpsnoop"
)

type HTTPReqInfo struct {
	Method    string
	Uri       string
	Referer   string
	Address   string
	Code      int
	Size      int64
	Duration  time.Duration
	UserAgent string
	Headers   map[string][]string `json:"-"`
}

func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

func requestGetRemoteAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		return parts[0]
	}
	return hdrRealIP
}

func logRequestHandler(h http.Handler, l logging.Logger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ri := &HTTPReqInfo{
			Method:    r.Method,
			Uri:       r.URL.String(),
			Referer:   r.Header.Get("Referer"),
			UserAgent: r.Header.Get("User-Agent"),
			Headers:   r.Header,
		}

		ri.Address = requestGetRemoteAddress(r)

		m := httpsnoop.CaptureMetrics(h, w, r)

		ri.Code = m.Code
		ri.Duration = m.Duration
		ri.Size = m.Written

		json, err := util.ToJson(ri)
		l.Debug("%s %s", json, err)
	}
	return http.HandlerFunc(fn)
}
