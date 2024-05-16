package main

/*

go run cmd/build-countries-js/main.go -metafile ~/Downloads/whosonfirst-data-country-latest/meta/whosonfirst-data-country-latest.csv > static/javascript/whosonfirst.spelunker.countries.js

*/

import (
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sfomuseum/go-csvdict"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd/templates/javascript"
)

type Country struct {
	WhosOnFirstId int64  `json:"wof:id"`
	Name          string `json:"wof:name"`
}

type TemplateVars struct {
	Lookup  string
	Created time.Time
}

func main() {

	var meta string

	flag.StringVar(&meta, "metafile", "", "The path to the country 'meta' CSV file.")

	flag.Parse()

	ctx := context.Background()

	t, err := javascript.LoadTemplates(ctx)

	if err != nil {
		log.Fatalf("Failed to load templates, %v", err)
	}

	countries_t := t.Lookup("whosonfirst_spelunker_countries")

	if countries_t == nil {
		log.Fatalf("Missing 'whosonfirst_spelunker_countries' template")
	}

	lookup := make(map[string]*Country)

	r, err := os.Open(meta)

	if err != nil {
		log.Fatalf("Failed to open %s for reading, %w", meta, err)
	}

	defer r.Close()

	csv_r, err := csvdict.NewReader(r)

	if err != nil {
		log.Fatalf("Failed to create CSV reader for %s, %v", meta, err)
	}

	for {
		row, err := csv_r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to read line, %v", err)
		}

		code, exists := row["wof_country"]

		if !exists {
			log.Fatalf("Row is missing 'wof_country', %v", row)
		}

		name, exists := row["name"]

		if !exists {
			log.Fatalf("Row is missing 'name', %v", row)
		}

		str_id, exists := row["id"]

		if !exists {
			log.Fatalf("Row is missing 'id', %v", row)
		}

		wof_id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			log.Fatalf("Failed to parse string ID '%s', %v", str_id, err)
		}

		lookup[code] = &Country{
			WhosOnFirstId: wof_id,
			Name:          name,
		}
	}

	enc_lookup, err := json.Marshal(lookup)

	if err != nil {
		log.Fatalf("Failed to encode lookup, %v", err)
	}

	created := time.Now()

	vars := TemplateVars{
		Lookup:  string(enc_lookup),
		Created: created,
	}

	err = countries_t.Execute(os.Stdout, vars)

	if err != nil {
		log.Fatalf("Failed to render template, %v", err)
	}

}
