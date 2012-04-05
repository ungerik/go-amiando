package amiando

import "fmt"

type UserDataType string

const (
	UserDataString      UserDataType = "string"   // value is of type String.
	UserDataNumber      UserDataType = "number"   // value is of type Integer.
	UserDataDate        UserDataType = "date"     // value is of type Date.
	UserDataGender      UserDataType = "gender"   // value is of type Integer.
	UserDataEmail       UserDataType = "email"    // value is of type String.
	UserDataUrl         UserDataType = "url"      // value is of type String.
	UserDataBirthday    UserDataType = "birthday" // value is of type Date.
	UserDataAddress     UserDataType = "address"  // value is an object of type Address.
	UserDataPhone       UserDataType = "phone"    // value is of type String.
	UserDataZipCode     UserDataType = "zipCode"  // value is of type String.
	UserDataCountry     UserDataType = "country"  // value is of type Country. Country codes are defined by the ISO 3166-1-alpha-2 code standard
	UserDataBlog        UserDataType = "blog"     // value is of type String.
	UserDataCheckbox    UserDataType = "checkbox" // value is of type Bool
	UserDataRadiobutton UserDataType = "radio"    // value is of type String.
	UserDataDropdown    UserDataType = "dropdown" // value is of type String.
	UserDataTextArea    UserDataType = "textarea" // value is of type String.
	UserDataProduct     UserDataType = "product"  // value is of type String.
	UserDataPhoto       UserDataType = "photo"    // value is of type String (URL)
)

type UserData struct {
	Title string       `json:"title"`
	Type  UserDataType `json:"type"`
	Value interface{}  `json:"value"`
}

func (self *UserData) String() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self *UserData) Address() *Address {
	if self.Type != UserDataAddress {
		return nil
	}
	data := self.Value.(map[string]interface{})
	addr := new(Address)
	if v, ok := data["street"]; ok {
		addr.Street = v.(string)
	}
	if v, ok := data["streets"]; ok {
		addr.Streets = v.(string)
	}
	if v, ok := data["city"]; ok {
		addr.City = v.(string)
	}
	if v, ok := data["zipCode"]; ok {
		addr.ZipCode = v.(string)
	}
	if v, ok := data["country"]; ok {
		addr.Country = v.(string)
	}
	return addr
}

type Address struct {
	Street  string
	Street2 string
	City    string
	ZipCode string
	Country string
}
