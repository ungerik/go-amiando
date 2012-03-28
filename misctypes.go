package amiando

import (
	"encoding/json"
	"strconv"
	"strings"
)

type PaymentStatus string
type UserDataType string
type TicketType string

///////////////////////////////////////////////////////////////////////////////
// ID

type ID int64

func (self ID) String() string {
	return strconv.FormatInt(int64(self), 10)
}

///////////////////////////////////////////////////////////////////////////////
// Error

type Error struct {
	errors []string
}

func (self *Error) Error() string {
	return strings.Join(self.errors, ", ")
}

///////////////////////////////////////////////////////////////////////////////
// ErrorReporter

type ErrorReporter interface {
	Err() error
	Reset()
}

///////////////////////////////////////////////////////////////////////////////
// ResultBase

type ResultBase struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

func (self *ResultBase) Err() error {
	if self.Success || len(self.Errors) == 0 {
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
