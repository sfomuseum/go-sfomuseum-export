package export

import (
	"encoding/json"
	"github.com/aaronland/go-artisanal-integers"
	brooklyn_integers "github.com/aaronland/go-brooklynintegers-api"
	wof_export "github.com/whosonfirst/go-whosonfirst-export"	
)

type Feature struct {
	Type       string      `json:"type"`
	Id         int64       `json:"id"`
	Properties interface{} `json:"properties"`
	Bbox       interface{} `json:"bbox,omitempty"`
	Geometry   interface{} `json:"geometry"`
}

type ExportOptions struct {
	Strict        bool
	IntegerClient artisanalinteger.Client
}

func DefaultExportOptions() (*ExportOptions, error) {

	bi_client := brooklyn_integers.NewAPIClient()

	opts := ExportOptions{
		Strict:        true,
		IntegerClient: bi_client,
	}

	return &opts, nil
}

type FeatureExporter struct {
	wof_export.Exporter
	options *ExportOptions
}

func NewExporter(opts *ExportOptions) (wof_export.Exporter, error) {

	ex := FeatureExporter{
		options: opts,
	}

	return &ex, nil
}

func (ex *FeatureExporter) ExportFeature(feature interface{}) ([]byte, error) {

	body, err := json.Marshal(feature)

	if err != nil {
		return nil, err
	}

	return ex.Export(body)
}

func (ex *FeatureExporter) Export(feature []byte) ([]byte, error) {

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

func Prepare(feature []byte, opts *ExportOptions) ([]byte, error) {

	var err error
	
	feature, err = wof_export.Prepare(feature, opts)

	if err != nil {
		return nil, err
	}

	// SFO Museum properties go here
	
	return feature, nil
}
