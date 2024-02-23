package funcs

import (
	"testing"

	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd"
)

func TestReplaceAll(t *testing.T) {

	uris := httpd.DefaultURIs()
	
	v := ReplaceAll(uris.Descendants, "{id}", int64(136251273))

	if v != "/id/136251273/descendants" {
		t.Fatalf("Failed replacement")
	}
}

func TestURIForId(t *testing.T) {

	uris := httpd.DefaultURIs()
	
	v := URIForId(uris.Descendants, int64(136251273))

	if v != "/id/136251273/descendants" {
		t.Fatalf("Failed to derive URI for ID")
	}
}
