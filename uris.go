package httpd

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type URIs struct {
	// WWW/human-readable
	Id          string `json:"id"`
	Descendants string `json:"descendants"`
	Search      string `json:"search"`
	About string `json:"about"`
	
	// Static assets
	Static string `json:"static"`

	// API/machine-readable
	GeoJSON string `json:"geojson"`
	GeoJSONLD string `json:"geojsonld"`
	NavPlace string `json:"navplace"`
	Select string `json:"select"`
	SPR     string `json:"spr"`		
	SVG     string `json:"svg"`
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
