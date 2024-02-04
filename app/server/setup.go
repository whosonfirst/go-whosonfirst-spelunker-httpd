package server

import (
	"context"
	"fmt"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd/templates/html"
)

func setupCommon() {

	ctx := context.Background()
	var err error

	sp, err = spelunker.NewSpelunker(ctx, spelunker_uri)

	if err != nil {
		setupCommonError = fmt.Errorf("Failed to set up network, %w", err)
	}

	html_templates, err = html.LoadTemplates(ctx)

	if err != nil {
		setupCommonError = fmt.Errorf("Failed to load HTML templates, %w", err)
	}

}
