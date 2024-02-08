package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-server/handler"
	_ "github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
)

func Run(ctx context.Context, logger *slog.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) error {

	opts, err := RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to derive run options from flagset, %w", err)
	}

	return RunWithOptions(ctx, opts, logger)
}

func RunWithOptions(ctx context.Context, opts *RunOptions, logger *slog.Logger) error {

	slog.SetDefault(logger)

	// First create a local copy of RunOptions that can't be
	// modified after the fact. 'run_options' is defined in vars.go

	v, err := opts.Clone()

	if err != nil {
		return fmt.Errorf("Failed to create local run options, %w", err)
	}

	run_options = v

	// To do: Add/consult "is enabled" flags

	handlers := map[string]handler.RouteHandlerFunc{

		// WWW/human-readable
		run_options.URIs.Descendants: descendantsHandlerFunc,
		run_options.URIs.Id:          idHandlerFunc,
		run_options.URIs.Search:      searchHandlerFunc,
		run_options.URIs.About:      aboutHandlerFunc,		

		// Static assets
		run_options.URIs.Static: staticHandlerFunc,

		// API/machine-readable
		run_options.URIs.GeoJSON: geoJSONHandlerFunc,
		run_options.URIs.GeoJSONLD: geoJSONLDHandlerFunc,
		run_options.URIs.NavPlace: navPlaceHandlerFunc,
		run_options.URIs.Select: selectHandlerFunc,		
		run_options.URIs.SPR:     sprHandlerFunc,		
		run_options.URIs.SVG:     svgHandlerFunc,
	}

	go func() {
		for uri, h := range handlers {
			slog.Info("Enable handler", "uri", uri, "handler", fmt.Sprintf("%T", h))
		}
	}()
	
        log_logger := slog.NewLogLogger(logger.Handler(), slog.LevelInfo)

        route_handler_opts := &handler.RouteHandlerOptions{
                           Handlers: handlers,
                           Logger: log_logger,
        }

        route_handler, err := handler.RouteHandlerWithOptions(route_handler_opts)

	if err != nil {
		return fmt.Errorf("Failed to configure route handler, %w", err)
	}
	
	mux := http.NewServeMux()
	mux.Handle("/", route_handler)

	s, err := server.NewServer(ctx, run_options.ServerURI)

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
