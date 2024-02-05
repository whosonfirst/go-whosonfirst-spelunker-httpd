package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd/www"
)

func descendantsHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		slog.Error("Failed to set up common configuration", "error", setupCommonError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupCommonError)
	}

	opts := &www.DescendantsHandlerOptions{
		Spelunker: sp,
		Templates: html_templates,
		URIs:      uris_table,
	}

	return www.DescendantsHandler(opts)
}

func idHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		slog.Error("Failed to set up common configuration", "error", setupCommonError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupCommonError)
	}

	opts := &www.IdHandlerOptions{
		Spelunker: sp,
		Templates: html_templates,
		URIs:      uris_table,
	}

	return www.IdHandler(opts)
}

func searchHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		slog.Error("Failed to set up common configuration", "error", setupCommonError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupCommonError)
	}

	opts := &www.SearchHandlerOptions{
		Spelunker: sp,
		Templates: html_templates,
		URIs:      uris_table,
	}

	return www.SearchHandler(opts)
}
