package api

import (
	"encoding/json"
	"io"
)

func parseBody(bdy io.ReadCloser, s interface{}) error {
	return json.NewDecoder(bdy).Decode(&s)
}
