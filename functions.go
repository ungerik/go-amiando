package amiando

import (
	"bytes"
	"encoding/json"
)

func PrettifyJSON(compactJSON []byte) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, compactJSON, "", "\t")
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
