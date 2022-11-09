package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/mikan/iso3166ja"
)

func main() {
	format := flag.String("f", "csv", "format (csv, json-array or json-map)")
	outputFileName := flag.String("o", "iso3166ja", "file name without extension")
	columns := flag.String("c", strings.Join(iso3166ja.DefaultCSVColumns, ","), "csv columns")
	flag.Parse()
	wikiText, err := iso3166ja.LoadWikipediaJA()
	if err != nil {
		log.Fatal(err)
	}
	countries, err := iso3166ja.Parse(wikiText)
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	var ext string
	switch strings.ToLower(*format) {
	case "csv":
		ext = "csv"
		data, err = iso3166ja.FormatCSV(countries, strings.Split(*columns, ",")...)
	case "json":
		fallthrough
	case "json-array":
		ext = "json"
		data, err = iso3166ja.FormatArrayJSON(countries)
	case "json-map":
		ext = "json"
		data, err = iso3166ja.FormatMapJSON(countries)
	default:
		println("unsupported format:", *format)
		os.Exit(2)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile(*outputFileName+"."+ext, data, 0644); err != nil {
		log.Fatal(err)
	}
}
