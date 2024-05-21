package main

// Meta file can be found here:
// https://data.geocode.earth/wof/dist/legacy/whosonfirst-data-country-latest.tar.bz2

// go run cmd/build-countries-go/main.go -metafile ~/Downloads/whosonfirst-data-country-latest/meta/whosonfirst-data-country-latest.csv > countries.go

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sfomuseum/go-csvdict"
	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd/templates/golang"
)

type Country struct {
	WhosOnFirstId int64  `json:"wof:id"`
	Name          string `json:"wof:name"`
}

type TemplateVars struct {
	Lookup    map[string]*Country
	Created   time.Time
	CreatedBy string
}

func main() {

	var meta string

	flag.StringVar(&meta, "metafile", "", "The path to the country 'meta' CSV file.")

	flag.Parse()

	ctx := context.Background()

	t, err := golang.LoadTemplates(ctx)

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

	created := time.Now()

	// There are all kinds of ways to do this but because we might just
	// be running this from "go -run" it all starts to get fiddly and
	// complicated and kind of a waste of time. So just be explicit.

	created_by := "build-countries-golang"

	vars := TemplateVars{
		Lookup:    lookup,
		Created:   created,
		CreatedBy: created_by,
	}

	err = countries_t.Execute(os.Stdout, vars)

	if err != nil {
		log.Fatalf("Failed to render template, %v", err)
	}

}
