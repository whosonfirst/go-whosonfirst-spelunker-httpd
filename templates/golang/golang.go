package golang

import (
	"context"
	"embed"

	sfomuseum_text "github.com/sfomuseum/go-template/text"
	"text/template"
)

//go:embed *.golang
var FS embed.FS

func LoadTemplates(ctx context.Context) (*template.Template, error) {

	return sfomuseum_text.LoadTemplatesMatching(ctx, "*.golang", FS)
}
