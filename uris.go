package httpd

import (
	"fmt"
	"log/slog"
	"net/url"
	"reflect"
	"strings"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
)

type URIs struct {
	// WWW/human-readable
	Id                     string   `json:"id"`
	IdAlt                  []string `json:"id_alt"`
	Concordances           string   `json:"concordances"`
	ConcordanceNS          string   `json:"concordance_ns"`
	ConcordanceNSPred      string   `json:"concordance_ns_pred"`
	ConcordanceNSPredValue string   `json:"concordance_ns_pred_value"`
	Descendants            string   `json:"descendants"`
	DescendantsAlt         []string `json:"descendants_alt"`
	Index                  string   `json:"index"`
	Placetypes             string   `json:"placetypes"`
	Placetype              string   `json:"placetype"`
	Recent                 string   `json:"recent"`
	RecentAlt              []string `json:"recent_alt"`
	Search                 string   `json:"search"`
	About                  string   `json:"about"`

	// Static assets
	Static string `json:"static"`

	// API/machine-readable
	DescendantsFaceted string   `json:"descendants_faceted"`
	GeoJSON            string   `json:"geojson"`
	GeoJSONAlt         []string `json:"geojson_alt"`
	GeoJSONLD          string   `json:"geojsonld"`
	GeoJSONLDAlt       []string `json:"geojsonld_alt"`
	NavPlace           string   `json:"navplace"`
	NavPlaceAlt        []string `json:"navplace_alt"`
	PlacetypeFaceted   string   `json:"placetype_faceted"`
	RecentFaceted      string   `json:"recent_faceted"`
	SearchFaceted      string   `json:"search_faceted"`
	Select             string   `json:"select"`
	SelectAlt          []string `json:"select_alt"`
	SPR                string   `json:"spr"`
	SPRAlt             []string `json:"spr_alt"`
	SVG                string   `json:"svg"`
	SVGAlt             []string `json:"svg_alt"`
}

func (u *URIs) ApplyPrefix(prefix string) error {

	val := reflect.ValueOf(*u)

	for i := 0; i < val.NumField(); i++ {

		field := val.Field(i)
		v := field.String()

		if v == "" {
			continue
		}

		if strings.HasPrefix(v, prefix) {
			continue
		}

		new_v, err := url.JoinPath(prefix, v)

		if err != nil {
			return fmt.Errorf("Failed to assign prefix to %s, %w", v, err)
		}

		reflect.ValueOf(u).Elem().Field(i).SetString(new_v)
	}

	return nil
}

func DefaultURIs() *URIs {

	uris_table := &URIs{

		// WWW/human-readable

		Index:                  "/",
		Search:                 "/search",
		About:                  "/about",
		Placetypes:             "/placetypes",
		Placetype:              "/placetypes/{placetype}",
		Concordances:           "/concordances/",
		ConcordanceNS:          "/concordances/{namespace}",
		ConcordanceNSPred:      "/concordances/{namespace}:{predicate}",
		ConcordanceNSPredValue: "/concordances/{namespace}:{predicate}={value}",
		Recent:                 "/recent/{duration}",
		RecentAlt: []string{
			"/recent",
		},
		Id:          "/id/{id}",
		Descendants: "/id/{id}/descendants",

		// Static Assets
		Static: "/static/",

		// API/machine-readable
		DescendantsFaceted: "/id/{id}/descendants/facets",

		GeoJSON: "/geojson/",
		GeoJSONAlt: []string{
			"/id/{id}/geojson",
		},
		GeoJSONLD: "/geojsonld/",
		GeoJSONLDAlt: []string{
			"/id/{id}/geojsonld",
		},
		NavPlace: "/navplace/",
		NavPlaceAlt: []string{
			"/id/{id}/navplace",
		},
		PlacetypeFaceted: "/placetypes/{placetype}/facets",
		RecentFaceted:    "/recent/{duration}/facets",
		Select:           "/select/",
		SelectAlt: []string{
			"/id/{id}/select",
		},
		SPR: "/spr/",
		SPRAlt: []string{
			"/id/{id}/spr",
		},
		SVG: "/svg/",
		SVGAlt: []string{
			"/id/{id}/svg",
		},
	}

	return uris_table
}

func URIForIdSimple(uri string, id int64) string {
	id_uri := ReplaceAll(uri, "{id}", id)
	return uriWithFilters(id_uri, nil, nil)
}

func URIForId(uri string, id int64, filters []spelunker.Filter, facets []spelunker.Facet) string {

	id_uri := ReplaceAll(uri, "{id}", id)
	return uriWithFilters(id_uri, filters, facets)
}

func URIForPlacetype(uri string, pt string, filters []spelunker.Filter, facets []spelunker.Facet) string {

	pt_uri := ReplaceAll(uri, "{placetype}", pt)
	return uriWithFilters(pt_uri, filters, facets)
}

func URIForRecent(uri string, d string, filters []spelunker.Filter, facets []spelunker.Facet) string {

	r_uri := ReplaceAll(uri, "{duration}", d)
	return uriWithFilters(r_uri, filters, facets)
}

func uriWithFilters(uri string, filters []spelunker.Filter, facets []spelunker.Facet) string {

	u, _ := url.Parse(uri)
	v := &url.Values{}

	for _, f := range filters {
		v.Set(f.Scheme(), fmt.Sprintf("%v", f.Value()))
	}

	for _, f := range facets {
		v.Set("facet", f.String())
	}

	u.RawQuery = v.Encode()

	slog.Debug("URI", "with filters and facets", u.String())
	return u.String()
}

func ReplaceAll(input string, pattern string, value any) string {
	str_value := fmt.Sprintf("%v", value)
	return strings.Replace(input, pattern, str_value, -1)
}
