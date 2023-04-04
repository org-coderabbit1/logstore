package httpreq

import (
	"context"
	"net/http"
	"strings"

	"github.com/weaveworks/common/middleware"
)

type headerContextKey string

var (
	// LogstoreActorPathHeader is the name of the header e.g. used to enqueue requests in hierarchical queues.
	LogstoreActorPathHeader = "X-Logstore-Actor-Path"

	// LogstoreActorPathDelimiter is the delimiter used to serialise the hierarchy of the actor.
	LogstoreActorPathDelimiter = "|"
)

func PropagateHeadersMiddleware(headers ...string) middleware.Interface {
	return middleware.Func(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			for _, h := range headers {
				value := req.Header.Get(h)
				if value != "" {
					ctx := req.Context()
					ctx = context.WithValue(ctx, headerContextKey(h), value)
					req = req.WithContext(ctx)
				}
			}
			next.ServeHTTP(w, req)
		})
	})
}

func ExtractHeader(ctx context.Context, name string) string {
	s, _ := ctx.Value(headerContextKey(name)).(string)
	return s
}

func ExtractActorPath(ctx context.Context) []string {
	value := ExtractHeader(ctx, LogstoreActorPathHeader)
	if value == "" {
		return nil
	}
	return strings.Split(value, LogstoreActorPathDelimiter)
}
