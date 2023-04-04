package logstore

import (
	"fmt"
	"net/http"

	"example.com/acme/logstore/pkg/logql/syntax"
	serverutil "example.com/acme/logstore/pkg/util/server"
)

func formatQueryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expr, err := syntax.ParseExpr(r.FormValue("query"))
		if err != nil {
			serverutil.WriteError(err, w)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "%s", syntax.Prettify(expr))
	}
}
