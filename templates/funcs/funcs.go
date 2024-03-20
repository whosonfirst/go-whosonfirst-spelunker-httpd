package funcs

import (
	"fmt"
	"log/slog"
	"net/url"

	"github.com/whosonfirst/go-whosonfirst-sources"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func NameForSource(prefix string) string {

	src, err := sources.GetSourceByName(prefix)

	if err != nil {
		return prefix
	}

	return src.Fullname
}

func FormatNumber(i int64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", i)
}

func AppendPagination(uri string, k string, v any) string {

	u, err := url.Parse(uri)

	if err != nil {
		slog.Error("Failed to parse URI to append pagination", "uri", uri, "error", err)
		return "#"
	}

	q := u.Query()
	q.Set(k, fmt.Sprintf("%v", v))

	u.RawQuery = q.Encode()
	return u.String()
}
