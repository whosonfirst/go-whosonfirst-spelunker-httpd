package www

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/aaronland/go-pagination"
	"github.com/sfomuseum/go-http-auth"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
)

type HasAlternatePlacetypeHandlerOptions struct {
	Spelunker     spelunker.Spelunker
	Authenticator auth.Authenticator
	Templates     *template.Template
	URIs          *httpd.URIs
}

type HasAlternatePlacetypeHandlerVars struct {
	PageTitle          string
	URIs               *httpd.URIs
	AlternatePlacetype string
	Places             []spr.StandardPlacesResult
	Pagination         pagination.Results
	PaginationURL      string
	FacetsURL          string
	FacetsContextURL   string
	OpenGraph          *OpenGraph
}

func HasAlternatePlacetypeHandler(opts *HasAlternatePlacetypeHandlerOptions) (http.Handler, error) {

	t_name := "alternate_placetype"
	t := opts.Templates.Lookup(t_name)

	if t == nil {
		return nil, fmt.Errorf("Failed to locate '%s' template", t_name)
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := httpd.LoggerWithRequest(req, nil)

		req_pt := req.PathValue("placetype")

		alt_pt := req_pt

		logger = logger.With("request placetype", req_pt)

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

		r, pg_r, err := opts.Spelunker.HasAlternatePlacetype(ctx, pg_opts, alt_pt, filters)

		if err != nil {
			logger.Error("Failed to get records having placetype", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		pagination_url := httpd.URIForPlacetype(opts.URIs.AlternatePlacetype, alt_pt, filters, nil)

		// This is not ideal but I am not sure what is better yet...
		facets_url := httpd.URIForPlacetype(opts.URIs.AlternatePlacetypeFaceted, alt_pt, filters, nil)
		facets_context_url := req.URL.Path

		vars := HasAlternatePlacetypeHandlerVars{
			PageTitle:          alt_pt,
			URIs:               opts.URIs,
			AlternatePlacetype: alt_pt,
			Places:             r.Results(),
			Pagination:         pg_r,
			PaginationURL:      pagination_url,
			FacetsURL:          facets_url,
			FacetsContextURL:   facets_context_url,
		}

		og_label := alt_pt

		og_title := fmt.Sprintf(`Who's On First \"%s\" records`, alt_pt)
		og_desc := fmt.Sprintf("Who's On First records that are %s", og_label)

		vars.OpenGraph = &OpenGraph{
			Type:        "Article",
			SiteName:    "Who's On First Spelunker",
			Title:       og_title,
			Description: og_desc,
			Image:       "",
		}

		rsp.Header().Set("Content-Type", "text/html")

		err = t.Execute(rsp, vars)

		if err != nil {
			logger.Error("Failed to render template", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
		}

	}

	h := http.HandlerFunc(fn)
	return h, nil
}
