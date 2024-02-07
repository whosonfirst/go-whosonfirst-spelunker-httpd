package server

import (
	"context"
	"fmt"

	"github.com/sfomuseum/go-http-auth"
	sfom_html "github.com/sfomuseum/go-template/html"
	"github.com/whosonfirst/go-whosonfirst-spelunker"
)

func setupCommon() {

	ctx := context.Background()
	var err error

	sp, err = spelunker.NewSpelunker(ctx, run_options.SpelunkerURI)

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

	authenticator, err = auth.NewAuthenticator(ctx, run_options.AuthenticatorURI)

	if err != nil {
		setupWWWError = fmt.Errorf("Failed to create new authenticator, %w", err)
		return
	}

	// To do: custom funcs...

	html_templates, err = sfom_html.LoadTemplates(ctx, run_options.Templates...)

	if err != nil {
		setupWWWError = fmt.Errorf("Failed to load HTML templates, %w", err)
		return
	}

}
