package server

import (
	html_template "html/template"
	"sync"

	"github.com/whosonfirst/go-whosonfirst-spelunker"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
)

var sp spelunker.Spelunker

var uris_table *httpd.URIs

var html_templates *html_template.Template

var setupCommonOnce sync.Once
var setupCommonError error
