package server

import (
	"context"
	"flag"
	"fmt"
	io_fs "io/fs"

	"github.com/mitchellh/copystructure"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd/static"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd/templates/html"
)

type RunOptions struct {
	ServerURI        string      `json:"server_uri"`
	SpelunkerURI     string      `json:"spelunker_uri"`
	AuthenticatorURI string      `json:"authenticator_uri"`
	URIs             *httpd.URIs `json:"uris"`
	Templates        []io_fs.FS  `json:"templates,omitemtpy"`
	StaticAssets     io_fs.FS    `json:"static_assets,omitempty"`
}

func (o *RunOptions) Clone() (*RunOptions, error) {

	v, err := copystructure.Copy(o)

	if err != nil {
		return nil, fmt.Errorf("Failed to create local run options, %w", err)
	}

	new_opts := v.(*RunOptions)

	new_opts.Templates = o.Templates
	new_opts.StaticAssets = o.StaticAssets

	return new_opts, nil
}

func RunOptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "SPELUNKER")

	if err != nil {
		return nil, fmt.Errorf("Failed to assign flags from environment variables, %w", err)
	}

	uris_table = &httpd.URIs{

		// WWW/human-readable

		// I can't get this to work...
		// Descendants: "/id/{id}/descendants",

		Id:          "/id",
		Descendants: "/descendants/",
		Search:      "/search/",

		// Static Assets
		Static: "/static/",

		// API/machine-readable
		GeoJSON: "/geojson/",
		SVG:     "/svg/",
	}

	opts := &RunOptions{
		ServerURI:        server_uri,
		AuthenticatorURI: authenticator_uri,
		SpelunkerURI:     spelunker_uri,
		URIs:             uris_table,
		Templates:        []io_fs.FS{html.FS},
		StaticAssets:     static.FS,
	}

	return opts, nil
}
