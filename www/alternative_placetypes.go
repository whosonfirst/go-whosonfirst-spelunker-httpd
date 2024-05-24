package www

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/sfomuseum/go-http-auth"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
)

type AlternatePlacetypesHandlerOptions struct {
	Spelunker     spelunker.Spelunker
	Authenticator auth.Authenticator
	Templates     *template.Template
	URIs          *httpd.URIs
}

type AlternatePlacetypesHandlerVars struct {
	PageTitle string
	URIs      *httpd.URIs
	Facets    []*spelunker.FacetCount
	OpenGraph *OpenGraph
}

func AlternatePlacetypesHandler(opts *AlternatePlacetypesHandlerOptions) (http.Handler, error) {

	t_name := "alternative_placetypes"
	t := opts.Templates.Lookup(t_name)

	if t == nil {
		return nil, fmt.Errorf("Failed to locate '%s' template", t_name)
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := httpd.LoggerWithRequest(req, nil)

		faceting, err := opts.Spelunker.GetAlternatePlacetypes(ctx)

		if err != nil {
			logger.Error("Failed to get alternative placetypes", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		vars := AlternatePlacetypesHandlerVars{
			PageTitle: "Alternate Placetypes",
			URIs:      opts.URIs,
			Facets:    faceting.Results,
		}

		vars.OpenGraph = &OpenGraph{
			Type:        "Article",
			SiteName:    "Who's On First Spelunker",
			Title:       "Who's On First Alternate Placetypes",
			Description: "Who's On First records grouped by their alternate placetypes",
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
