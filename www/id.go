package www

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-pagination/countable"
	"github.com/sfomuseum/go-http-auth"
	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
)

type IdHandlerOptions struct {
	Spelunker     spelunker.Spelunker
	Authenticator auth.Authenticator
	Templates     *template.Template
	URIs          *httpd.URIs
}

type IdHandlerVars struct {
	Id               int64
	PageTitle        string
	URIs             *httpd.URIs
	Properties       string
	CountDescendants int64
}

func IdHandler(opts *IdHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("id")

	if t == nil {
		return nil, fmt.Errorf("Failed to locate 'id' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		logger := slog.Default()
		logger = logger.With("request", req.URL)

		uri, err, status := httpd.ParseURIFromRequest(req, nil)

		if err != nil {
			slog.Error("Failed to parse URI from request", "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), status)
			return
		}

		logger = logger.With("wofid", uri.Id)

		f, err := opts.Spelunker.GetById(ctx, uri.Id)

		if err != nil {
			slog.Error("Failed to get by ID", "id", uri.Id, "error", err)
			http.Error(rsp, spelunker.ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		pg_opts, err := countable.NewCountableOptions()

		if err != nil {
			slog.Error("Failed to create countable pagination options", "id", uri.Id, "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		pg_opts.PerPage(1)

		_, pg_rsp, err := opts.Spelunker.GetDescendants(ctx, uri.Id, pg_opts)

		if err != nil {
			slog.Error("Failed to count descendants", "id", uri.Id, "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		count_descendants := pg_rsp.Total()
		slog.Info("COUNT", "total", count_descendants)

		props := gjson.GetBytes(f, "properties")
		page_title := gjson.GetBytes(f, "properties.wof:name")

		vars := IdHandlerVars{
			Id:               uri.Id,
			Properties:       props.String(),
			PageTitle:        page_title.String(),
			URIs:             opts.URIs,
			CountDescendants: count_descendants,
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
