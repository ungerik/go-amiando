package amiando

import "fmt"

type Participant struct {
	Event *Event `json:"-"`

	PaymentID     ID            `json:"-"`
	PaymentUserID ID            `json:"buyerId"`      // payment
	PaymentStatus PaymentStatus `json:"status"`       // payment
	InvoiceNumber string        `json:"identifier"`   // payment
	CreatedDate   string        `json:"creationTime"` // payment
	ModifiedDate  string        `json:"lastModified"` // payment

	UserData []UserData `json:"userData"` // payment & ticket

	TicketID           ID         `json:"-"`
	FirstName          string     `json:"firstName"`         // ticket
	LastName           string     `json:"lastName"`          // ticket
	Email              string     `json:"email"`             // ticket
	CheckedDate        string     `json:"lastChecked"`       // ticket
	CancelledDate      string     `json:"cancelled"`         // ticket
	TicketType         TicketType `json:"ticketType"`        // ticket
	RegistrationNumber string     `json:"displayIdentifier"` // ticket
}

// Returns nil if no UserData with title is found
func (self *Participant) FindUserData(title string, restrictToTypes ...UserDataType) (userData *UserData, found bool) {
	for _, u := range self.UserData {
		if u.Title == title {
			if len(restrictToTypes) == 0 {
				return &u, true
			} else {
				for _, restrictToType := range restrictToTypes {
					if u.Type == restrictToType {
						return &u, true
					}
				}
			}
		}
	}
	return nil, false
}

func (self *Participant) FindRequiredUserData(title string, restrictToTypes ...UserDataType) (userData *UserData, err error) {
	userData, found := self.FindUserData(title, restrictToTypes...)
	if !found {
		return nil, fmt.Errorf("Required UserData \"%s\" of participant \"%s %s\"<%s> not found", title, self.FirstName, self.LastName, self.Email)
	}
	return userData, nil
}
