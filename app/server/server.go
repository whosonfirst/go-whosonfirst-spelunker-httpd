package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-server/handler"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
)

type RunOptions struct {
	Logger           *slog.Logger
	ServerURI        string
	SpelunkerURI     string
	AuthenticatorURI string
}

func Run(ctx context.Context, logger *slog.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) error {

	opts, err := RunOptionsFromFlagSet(ctx, fs, logger)

	if err != nil {
		return fmt.Errorf("Failed to derive run options from flagset, %w", err)
	}

	return RunWithOptions(ctx, opts)
}

func RunOptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) (*RunOptions, error) {

	flagset.Parse(fs)

	opts := &RunOptions{
		Logger:           logger,
		ServerURI:        server_uri,
		AuthenticatorURI: authenticator_uri,
		SpelunkerURI:     spelunker_uri,
	}

	return opts, nil
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	slog.SetDefault(opts.Logger)

	uris_table = &httpd.URIs{
		// WWW/human-readable
		Id:          "/id/",
		Descendants: "/descendants/", // FIX ME: Update to use improved syntax in Go 1.22
		Search:      "/search/",

		// API/machine-readable
		GeoJSON: "/geojson/",
		SVG:     "/svg/",
	}

	handlers := map[string]handler.RouteHandlerFunc{

		// WWW/human-readable
		uris_table.Descendants: descendantsHandlerFunc,
		uris_table.Id:          idHandlerFunc,
		uris_table.Search:      searchHandlerFunc,

		// API/machine-readable
		uris_table.GeoJSON: geoJSONHandlerFunc,
		uris_table.SVG:     svgHandlerFunc,
	}

	go func() {
		for uri, h := range handlers {
			slog.Debug("Enable handler", "uri", uri, "handler", fmt.Sprintf("%T", h))
		}
	}()

	route_handler, err := handler.RouteHandler(handlers)

	if err != nil {
		return fmt.Errorf("Failed to create route handlers, %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", route_handler)

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new server, %w", err)
	}

	slog.Info("Listening for requests", "address", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to start server, %w", err)
	}

	return nil
}
