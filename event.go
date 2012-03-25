package amiando

import "fmt"

///////////////////////////////////////////////////////////////////////////////
// Event

type BasicEventData struct {
	HostID                 ID      `json:"hostId"`
	Title                  string  `json:"title"`
	Country                string  `json:"country"`
	Language               string  `json:"language"`
	StartDate              string  `json:"selectedDate"`
	EndDate                string  `json:"selectedEndDate"`
	Timezone               string  `json:"timezone"`
	Visibility             string  `json:"visibility"`
	Identifier             string  `json:"identifier"`
	Description            string  `json:"description"`
	ShortDescription       string  `json:"shortDescription"`
	EventType              string  `json:"eventType"`
	OrganisatorDisplayName string  `json:"organisatorDisplayName"`
	PartnerEventUrl        string  `json:"partnerEventUrl"`
	Location               string  `json:"location"`
	LocationDescription    string  `json:"locationDescription"`
	Street                 string  `json:"street2"`
	ZipCode                string  `json:"zipCode"`
	City                   string  `json:"city"`
	State                  string  `json:"state"`
	CreationTime           string  `json:"creationTime"`
	LastModified           string  `json:"lastModified"`
	Longitude              float64 `json:"longitude"`
	Latitude               float64 `json:"latitude"`
}

type Event struct {
	ResultBase
	BasicEventData `json:"event"`
	Api            *Api
	Identifier     string
	InternalID     ID
}

func NewEvent(api *Api, identifier string) (event *Event, err error) {
	event = &Event{
		Api:        api,
		Identifier: identifier,
	}

	// Search for event with identifier
	type Result struct {
		ResultBase
		Ids []ID `json:"ids"`
	}
	var result Result
	err = event.Api.Call("event/find?identifier=%s", identifier, &result)
	if err != nil {
		return nil, err
	}
	if len(result.Ids) == 0 {
		return nil, fmt.Errorf("No event found for identifier '%s'", identifier)
	}
	// Find event with exact match of identifier
	// because API find returns all events whose identifiers include the searched one
	for _, id := range result.Ids {
		type Result struct {
			ResultBase
			BasicEventData `json:"event"`
		}
		var result Result
		err = event.Api.Call("event/%v", id, &result)
		if err != nil {
			return nil, err
		}
		if result.Identifier == identifier {
			event.InternalID = id
			break
		}
	}
	if event.InternalID == 0 {
		return nil, fmt.Errorf("No exact match found for identifier '%s'", identifier)
	}

	err = event.Read(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (self *Event) Read(out ErrorReporter) (err error) {
	return self.Api.Call("event/%v", self.InternalID, out)
}

func (self *Event) PaymentIDs() (ids []ID, err error) {
	type Result struct {
		ResultBase
		Payments []ID `json:"payments"`
	}
	var result Result
	err = self.Api.Call("event/%v/payments", self.InternalID, &result)
	if err != nil {
		return nil, err
	}
	return result.Payments, nil
}

func (self *Event) TicketIDs() (ids []ID, err error) {
	type Result struct {
		ResultBase
		Ids []ID `json:"ids"`
	}
	var result Result
	err = self.Api.Call("ticket/find?eventId=%v", self.InternalID, &result)
	if err != nil {
		return nil, err
	}
	return result.Ids, nil
}

func (self *Event) Participants() (participants []*Participant, err error) {
	participants = []*Participant{}

	paymentIDs, err := self.PaymentIDs()
	if err != nil {
		return nil, err
	}

	for _, paymentID := range paymentIDs {
		ticketIDs, err := self.Api.TicketIDsOfPayment(paymentID)
		if err != nil {
			return nil, err
		}

		for i, ticketID := range ticketIDs {
			participant := &Participant{
				Event:     self,
				PaymentID: paymentID,
				TicketID:  ticketID,
			}

			err = self.Api.Payment(paymentID, participant)
			if err != nil {
				return nil, err
			}

			// Save payment UserData because it will be overwritten by the ticket UserData 
			userData := participant.UserData

			err = self.Api.Ticket(ticketID, participant)
			if err != nil {
				return nil, err
			}

			// If there is no ticket UserData use payment UserData for the first ticket
			if i == 0 && len(participant.UserData) == 0 {
				participant.UserData = userData
			}

			participants = append(participants, participant)
		}
	}

	return participants, nil
}

func (self *Event) EnumParticipants() (<-chan *Participant, <-chan error) {
	p := make(chan *Participant, 32)
	e := make(chan error, 1)

	go func() {
		defer close(p)
		defer close(e)

		paymentIDs, err := self.PaymentIDs()
		if err != nil {
			e <- err
			return
		}
		for _, paymentID := range paymentIDs {
			ticketIDs, err := self.Api.TicketIDsOfPayment(paymentID)
			if err != nil {
				e <- err
				return
			}

			for i, ticketID := range ticketIDs {
				participant := &Participant{
					Event:     self,
					PaymentID: paymentID,
					TicketID:  ticketID,
				}

				err = self.Api.Payment(paymentID, participant)
				if err != nil {
					e <- err
					return
				}

				// Save payment UserData because it will be overwritten by the ticket UserData 
				userData := participant.UserData

				err = self.Api.Ticket(ticketID, participant)
				if err != nil {
					e <- err
					return
				}

				// If there is no ticket UserData use payment UserData for the first ticket
				if i == 0 && len(participant.UserData) == 0 {
					participant.UserData = userData
				}

				p <- participant
			}
		}
	}()

	return p, e
}
