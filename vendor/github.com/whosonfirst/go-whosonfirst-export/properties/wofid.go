package properties

import (
	"github.com/aaronland/go-artisanal-integers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func EnsureWOFId(feature []byte, client artisanalinteger.Client) ([]byte, error) {

	var err error

	var wof_id int64

	rsp := gjson.GetBytes(feature, "properties.wof:id")

	if rsp.Exists() {

		wof_id = rsp.Int()

	} else {

		i, err := client.NextInt()

		if err != nil {
			return nil, err
		}

		wof_id = i

		feature, err = sjson.SetBytes(feature, "properties.wof:id", wof_id)

		if err != nil {
			return nil, err
		}
	}

	id := gjson.GetBytes(feature, "id")

	if !id.Exists() {

		feature, err = sjson.SetBytes(feature, "id", wof_id)

		if err != nil {
			return nil, err
		}

	}

	return feature, nil
}
