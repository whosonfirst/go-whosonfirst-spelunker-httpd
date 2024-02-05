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

func Run(ctx context.Context, logger *slog.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) error {

	flagset.Parse(fs)

	slog.SetDefault(logger)

	uris_table = &httpd.URIs{
		// WWW/human-readable
		Id:          "/id/",
		Descendants: "/descendants/", // FIX ME: Update to use improved syntax in Go 1.22
		Search:      "/search/",

		// API/machine-readable
		GeoJSON: "/geojson",
	}

	handlers := map[string]handler.RouteHandlerFunc{

		// WWW/human-readable
		uris_table.Descendants: descendantsHandlerFunc,
		uris_table.Id:          idHandlerFunc,
		uris_table.Search:      searchHandlerFunc,

		// API/machine-readable
		uris_table.GeoJSON: geoJSONHandlerFunc,
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

	logger.Info("Listening for requests", "address", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to start server, %w", err)
	}

	return nil
}
