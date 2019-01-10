package export

import (
	"encoding/json"
	wof_export "github.com/whosonfirst/go-whosonfirst-export"
	wof_exporter "github.com/whosonfirst/go-whosonfirst-export/exporter"
	wof_options "github.com/whosonfirst/go-whosonfirst-export/options"
)

type SFOMuseumExporter struct {
	wof_exporter.Exporter
	options wof_options.Options
}

func NewSFOMuseumExporter(opts wof_options.Options) (wof_exporter.Exporter, error) {

	ex := SFOMuseumExporter{
		options: opts,
	}

	return &ex, nil
}

func (ex *SFOMuseumExporter) ExportFeature(feature interface{}) ([]byte, error) {

	body, err := json.Marshal(feature)

	if err != nil {
		return nil, err
	}

	return ex.Export(body)
}

func (ex *SFOMuseumExporter) Export(feature []byte) ([]byte, error) {

	var err error

	feature, err = Prepare(feature, ex.options)

	if err != nil {
		return nil, err
	}

	feature, err = wof_export.Format(feature, ex.options)

	if err != nil {
		return nil, err
	}

	return feature, nil
}

func Prepare(feature []byte, opts wof_options.Options) ([]byte, error) {

	var err error

	feature, err = wof_export.Prepare(feature, opts)

	if err != nil {
		return nil, err
	}

	// SFO Museum properties go here

	return feature, nil
}
