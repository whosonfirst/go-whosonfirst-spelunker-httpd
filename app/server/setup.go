package server

import (
	"context"
	"fmt"

	"github.com/sfomuseum/go-http-auth"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd/templates/html"
)

func setupCommon() {

	ctx := context.Background()
	var err error

	sp, err = spelunker.NewSpelunker(ctx, spelunker_uri)

	if err != nil {
		setupCommonError = fmt.Errorf("Failed to set up network, %w", err)
		return
	}
}

func setupWWW() {

	ctx := context.Background()
	var err error

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		setupWWWError = fmt.Errorf("Common setup failed, %w", err)
		return
	}

	authenticator, err = auth.NewAuthenticator(ctx, authenticator_uri)

	if err != nil {
		setupWWWError = fmt.Errorf("Failed to create new authenticator, %w", err)
		return
	}

	html_templates, err = html.LoadTemplates(ctx)

	if err != nil {
		setupWWWError = fmt.Errorf("Failed to load HTML templates, %w", err)
		return
	}

}
