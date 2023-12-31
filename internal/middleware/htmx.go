package middleware

import (
	"context"
	"net/http"
)

var htmxRequesHeadersContextKey contextKey = "htmx-request-headers"

// https://htmx.org/reference/#request_headers
type HtmxRequestHeaders struct {
	HXBoosted               string
	HXCurrentURL            string
	HXHistoryRestoreRequest string
	HXPrompt                string
	HXRequest               string
	HXTarget                string
	HXTriggerName           string
	HXTrigger               string
}

func (h HtmxRequestHeaders) IsBoosted() bool {
	return h.HXBoosted == "true"
}

func HtmxRequestHeadersFromContext(ctx context.Context) HtmxRequestHeaders {
	if htmx, ok := ctx.Value(htmxRequesHeadersContextKey).(HtmxRequestHeaders); ok {
		return htmx
	}
	return HtmxRequestHeaders{}
}

func CaptureHtmxRequestHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var headers HtmxRequestHeaders
		headers.HXBoosted = r.Header.Get("HX-Boosted")
		headers.HXCurrentURL = r.Header.Get("HX-Current-URL")
		headers.HXHistoryRestoreRequest = r.Header.Get("HX-History-Restore-Request")
		headers.HXPrompt = r.Header.Get("HX-Prompt")
		headers.HXRequest = r.Header.Get("HX-Request")
		headers.HXTarget = r.Header.Get("HX-Target")
		headers.HXTriggerName = r.Header.Get("HX-Trigger-Name")
		headers.HXTrigger = r.Header.Get("HX-Trigger")

		r = r.WithContext(context.WithValue(r.Context(), htmxRequesHeadersContextKey, headers))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
