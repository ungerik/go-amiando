package amiando

import (
	"os"
	"json"
	"bytes"
	"strings"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// ID

type ID int64

func (self ID) String() string {
	return strconv.Itoa64(int64(self))
}

///////////////////////////////////////////////////////////////////////////////
// Error

type Error struct {
	errors []string
}

func (self *Error) String() string {
	return strings.Join(self.errors, ", ")
}

///////////////////////////////////////////////////////////////////////////////
// ErrorReporter

type ErrorReporter interface {
	Error() os.Error
	Reset()
}

///////////////////////////////////////////////////////////////////////////////
// ResultBase

type ResultBase struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

func (self *ResultBase) Error() (err os.Error) {
	if self.Success {
		return nil
	}
	return &Error{self.Errors}
}

func (self *ResultBase) Reset() {
	self.Success = false
	self.Errors = nil
}

///////////////////////////////////////////////////////////////////////////////
// JsonResult

// Has to be used as a pointer to member of a struct
type JsonResult struct {
	ResultBase
	JSON []byte
}

func (self *JsonResult) UnmarshalJSON(jsonData []byte) os.Error {
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

///////////////////////////////////////////////////////////////////////////////
// Functions

func PrettifyJSON(compactJSON []byte) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, compactJSON, "", "\t")
	if err != nil {
		return err.String()
	}
	return buf.String()
}
