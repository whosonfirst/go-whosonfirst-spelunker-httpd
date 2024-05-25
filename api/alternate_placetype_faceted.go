package api

import (
	"encoding/json"
	"net/http"

	// TBD...
	// "github.com/sfomuseum/go-http-auth"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
)

type AlternatePlacetypeFacetedHandlerOptions struct {
	Spelunker spelunker.Spelunker
	// TBD...
	// Authenticator auth.Authenticator
}

func AlternatePlacetypeFacetedHandler(opts *AlternatePlacetypeFacetedHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := httpd.LoggerWithRequest(req, nil)

		req_pt := req.PathValue("placetype")
		logger = logger.With("request placetype", req_pt)

		alt_pt := req_pt

		filter_params := httpd.DefaultFilterParams()

		filters, err := httpd.FiltersFromRequest(ctx, req, filter_params)

		if err != nil {
			logger.Error("Failed to derive filters from request", "error", err)
			http.Error(rsp, "Bad request", http.StatusBadRequest)
			return
		}

		facets, err := httpd.FacetsFromRequest(ctx, req, filter_params)

		if err != nil {
			logger.Error("Failed to derive facets from requrst", "error", err)
			http.Error(rsp, "Bad request", http.StatusBadRequest)
			return
		}

		if len(facets) == 0 {
			logger.Error("No facets from requrst")
			http.Error(rsp, "Bad request", http.StatusBadRequest)
			return
		}

		facets_rsp, err := opts.Spelunker.HasAlternatePlacetypeFaceted(ctx, alt_pt, filters, facets)

		if err != nil {
			logger.Error("Failed to get facets for alternate placetype", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")

		enc := json.NewEncoder(rsp)
		err = enc.Encode(facets_rsp)

		if err != nil {
			logger.Error("Failed to encode facets response", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

	}

	h := http.HandlerFunc(fn)
	return h, nil
}
