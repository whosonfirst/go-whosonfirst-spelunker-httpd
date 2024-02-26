package funcs

import (
	"fmt"
	"strings"
)

// Custom template functions

func URIForId(uri string, id int64) string {
	return ReplaceAll(uri, "{id}", id)
}

func ReplaceAll(input string, pattern string, value any) string {
	str_value := fmt.Sprintf("%v", value)
	return strings.Replace(input, pattern, str_value, -1)
}
