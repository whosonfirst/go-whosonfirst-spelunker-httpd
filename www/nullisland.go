package www

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-pagination"
	"github.com/sfomuseum/go-http-auth"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
)

type NullIslandHandlerOptions struct {
	Spelunker     spelunker.Spelunker
	Authenticator auth.Authenticator
	Templates     *template.Template
	URIs          *httpd.URIs
}

type NullIslandHandlerVars struct {
	PageTitle        string
	URIs             *httpd.URIs
	Places           []spr.StandardPlacesResult
	Pagination       pagination.Results
	PaginationURL    string
	FacetsURL        string
	FacetsContextURL string
}

func NullIslandHandler(opts *NullIslandHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("nullisland")

	if t == nil {
		return nil, fmt.Errorf("Failed to locate 'recent' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		logger := slog.Default()
		logger = logger.With("request", req.URL)

		pg_opts, err := httpd.PaginationOptionsFromRequest(req)

		if err != nil {
			logger.Error("Failed to create pagination options", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		filter_params := httpd.DefaultFilterParams()

		filters, err := httpd.FiltersFromRequest(ctx, req, filter_params)

		if err != nil {
			logger.Error("Failed to derive filters from request", "error", err)
			http.Error(rsp, "Bad request", http.StatusBadRequest)
			return
		}

		r, pg_r, err := opts.Spelunker.VisitingNullIsland(ctx, pg_opts, filters)

		if err != nil {
			logger.Error("Failed to get recent", "error", err)
			http.Error(rsp, "InternalServerError", http.StatusInternalServerError)
			return
		}

		// This is not ideal but I am not sure what is better yet...
		pagination_url := httpd.URIForNullIsland(opts.URIs.NullIsland, filters, nil)

		// This is not ideal but I am not sure what is better yet...
		facets_url := httpd.URIForNullIsland(opts.URIs.NullIslandFaceted, filters, nil)
		facets_context_url := req.URL.Path

		vars := NullIslandHandlerVars{
			Places:           r.Results(),
			Pagination:       pg_r,
			URIs:             opts.URIs,
			PaginationURL:    pagination_url,
			FacetsURL:        facets_url,
			FacetsContextURL: facets_context_url,
		}

		rsp.Header().Set("Content-Type", "text/html")

		err = t.Execute(rsp, vars)

		if err != nil {
			logger.Error("Failed to return ", "error", err)
			http.Error(rsp, "InternalServerError", http.StatusInternalServerError)
		}

	}

	h := http.HandlerFunc(fn)
	return h, nil
}
