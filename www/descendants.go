package www

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-pagination/countable"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
)

type DescendantsHandlerOptions struct {
	Spelunker spelunker.Spelunker
	Templates *template.Template
}

type DescendantsHandlerVars struct {
	Places     []spr.StandardPlacesResult
	Pagination pagination.Results
}

func DescendantsHandler(opts *DescendantsHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("descendants")

	if t == nil {
		return nil, fmt.Errorf("Failed to locate 'descendants' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		logger := slog.Default()
		logger = logger.With("request", req.URL)

		slog.Info("Get descendants")
		
		uri, err, status := httpd.ParseURIFromRequest(req, nil)

		if err != nil {
			slog.Error("Failed to parse URI from request", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), status)
			return
		}

		logger = logger.With("wofid", uri.Id)

		pg_opts, err := countable.NewCountableOptions()

		if err != nil {
			slog.Error("Failed to create pagination options", "error", err)
			http.Error(rsp, "womp womp", http.StatusInternalServerError)
			return
		}

		r, pg_r, err := opts.Spelunker.GetDescendants(ctx, uri.Id, pg_opts)

		if err != nil {
			slog.Error("Failed to get descendants", "error", err)
			http.Error(rsp, "womp womp", http.StatusInternalServerError)
			return
		}

		vars := DescendantsHandlerVars{
			Places:     r.Results(),
			Pagination: pg_r,
		}

		rsp.Header().Set("Content-Type", "text/html")

		err = t.Execute(rsp, vars)

		if err != nil {
			slog.Error("Failed to return ", "error", err)
			http.Error(rsp, "womp womp", http.StatusInternalServerError)
		}

	}

	h := http.HandlerFunc(fn)
	return h, nil
}
