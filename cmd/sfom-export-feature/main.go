package main

import (
	"bytes"
	"context"
	"flag"
	_ "github.com/sfomuseum/go-sfomuseum-export/v2"
	"github.com/whosonfirst/go-whosonfirst-export/v2"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	exporter_uri := flag.String("exporter-uri", "sfomuseum://", "A valid whosonfirst/go-whosonfirst-export URI")

	flag.Parse()

	ctx := context.Background()

	ex, err := export.NewExporter(ctx, *exporter_uri)

	if err != nil {
		log.Fatalf("Failed to create new exporter for '%s', %v", *exporter_uri, err)
	}

	writers := make([]io.Writer, 0)
	writers = append(writers, os.Stdout)
	wr := io.MultiWriter(writers...)

	for _, path := range flag.Args() {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatalf("Failed to open '%s', %v", path, err)
		}

		body, err := ioutil.ReadAll(fh)

		if err != nil {
			log.Fatalf("Failed to read '%s', %v", path, err)
		}

		pretty, err := ex.Export(ctx, body)

		if err != nil {
			log.Fatalf("Failed to export '%s', %v", path, err)
		}

		br := bytes.NewReader(pretty)
		_, err = io.Copy(wr, br)

		if err != nil {
			log.Fatalf("Failed to write feature, %v", err)
		}
	}

	os.Exit(0)
}
