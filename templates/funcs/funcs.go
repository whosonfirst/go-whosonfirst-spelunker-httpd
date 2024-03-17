package funcs

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/whosonfirst/go-whosonfirst-sources"
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
