package amiando

import (
	"fmt"
)

type UserData struct {
	Title string       `json:"title"`
	Type  UserDataType `json:"type"`
	Value interface{}  `json:"value"`
}

func (self *UserData) String() string {
	return fmt.Sprintf("%v", self.Value)
}
