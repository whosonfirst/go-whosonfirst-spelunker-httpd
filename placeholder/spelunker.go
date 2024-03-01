package placeholder

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/aaronland/go-pagination"
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	wof_spr "github.com/whosonfirst/go-whosonfirst-spr/v2"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"github.com/sfomuseum/go-placeholder-client"
)

type PlaceholderSpelunker struct {
	spelunker.Spelunker
	client *client.PlaceholderClient
}

func init() {
	ctx := context.Background()
	spelunker.RegisterSpelunker(ctx, "placeholder", NewPlaceholderSpelunker)
}

func NewPlaceholderSpelunker(ctx context.Context, uri string) (spelunker.Spelunker, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	endpoint := q.Get("endpoint")

	if endpoint == "" {
		return nil, fmt.Errorf("Missing ?endpoint= parameter")
	}

	cl, err := client.NewPlaceholderClient(endpoint)
	
	if err != nil {
		return nil, fmt.Errorf("Failed to create placeholder client, %w", err)
	}

	s := &PlaceholderSpelunker{
		client: cl,
	}

	return s, nil
}

func (s *PlaceholderSpelunker) Search(ctx context.Context, pg_opts pagination.Options, search_opts *spelunker.SearchOptions) (wof_spr.StandardPlacesResults, pagination.Results, error) {

	results, err := s.client.Search(search_opts.Query)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to perform query, %w", err)
	}

	slog.Info("R", "r", results)

	return nil, nil, spelunker.ErrNotImplemented
}


func (s *PlaceholderSpelunker) GetById(ctx context.Context, id int64) ([]byte, error) {

	return nil, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) GetAlternateGeometryById(ctx context.Context, id int64, alt_geom *uri.AltGeom) ([]byte, error) {

	return nil, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) getById(ctx context.Context, q string, args ...interface{}) ([]byte, error) {

	return nil, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) GetDescendants(ctx context.Context, pg_opts pagination.Options, id int64, filters []spelunker.Filter) (wof_spr.StandardPlacesResults, pagination.Results, error) {

	return nil, nil, spelunker.ErrNotImplemented	
}

func (s *PlaceholderSpelunker) GetDescendantsFaceted(ctx context.Context, id int64, filters []spelunker.Filter, facets []*spelunker.Facet) ([]*spelunker.Faceting, error) {

	return nil, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) CountDescendants(ctx context.Context, id int64) (int64, error) {

	return 0, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) HasPlacetype(ctx context.Context, pg_opts pagination.Options, pt *placetypes.WOFPlacetype, filters []spelunker.Filter) (wof_spr.StandardPlacesResults, pagination.Results, error) {

	return nil, nil, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) GetRecent(ctx context.Context, pg_opts pagination.Options, d time.Duration, filters []spelunker.Filter) (wof_spr.StandardPlacesResults, pagination.Results, error) {

	return nil, nil, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) GetPlacetypes(ctx context.Context) (*spelunker.Faceting, error) {

	return nil, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) GetConcordances(ctx context.Context) (*spelunker.Faceting, error) {

	return nil, spelunker.ErrNotImplemented
}

func (s *PlaceholderSpelunker) HasConcordance(ctx context.Context, pg_opts pagination.Options, namespace string, predicate string, value string, filters []spelunker.Filter) (wof_spr.StandardPlacesResults, pagination.Results, error) {
	
	return nil, nil, spelunker.ErrNotImplemented
}
