package amiando

const DateFormat = "2006-01-02T15:04:05"

type PaymentStatus string

const (
	PaymentNew        PaymentStatus = "new"
	PaymentAuthorized PaymentStatus = "authorized"
	PaymentPaid       PaymentStatus = "paid"
	PaymentDisbursed  PaymentStatus = "disbursed"
	PaymentCancelled  PaymentStatus = "cancelled"
)

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
)

type TicketType string

func (self TicketType) String() string {
	return string(self)
}

const (
	BadgeTicket        TicketType = "com.amiando.ticket.type.Badge"        // Means that the ticket is a badge.
	ETicketTicket      TicketType = "com.amiando.ticket.type.ETicket"      // Means that the ticket will be sent via email.
	PaperTicket        TicketType = "com.amiando.ticket.type.Paper"        // Means that the ticket is a confirmation.
	ConfirmationTicket TicketType = "com.amiando.ticket.type.Confirmation" // Means that the payment was bought using prepayment.
	OnSiteTicket       TicketType = "com.amiando.ticket.type.OnSite"       // Means that the ticket was bought via EasyEntry.
)

const (
	Male   = 1
	Female = 2
)
