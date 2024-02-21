package www

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aaronland/go-http-sanitize"
)

func FacetsFromRequest(ctx context.Context, req *http.Request, params ...string) ([]string, error) {

	// TBD...
	facets := make([]string, 0)
	
	v, err := sanitize.GetString(req, "facet")
	
	if err != nil {
		return nil, fmt.Errorf("Failed to derive ?facet= query  parameter, %w", err)
	}
	
	if v == "" {
		return nil, fmt.Errorf("Empty facet paramter")
	}
	
	facets = append(facets, v)
	return facets, nil
}
