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

///////////////////////////////////////////////////////////////////////////////
// JsonResult

// Has to be used as a pointer to member of a struct
type JsonResult struct {
	ResultBase
	JSON []byte
}

func (self *JsonResult) UnmarshalJSON(jsonData []byte) error {
	self.JSON = jsonData
	return json.Unmarshal(jsonData, &self.ResultBase)
}

func (self *JsonResult) String() string {
	return PrettifyJSON(self.JSON)
}

func (self *JsonResult) Reset() {
	self.ResultBase.Reset()
	self.JSON = nil
}
