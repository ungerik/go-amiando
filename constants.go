package amiando

const DateFormat = "2006-01-02T15:04:05"

const (
	PaymentNew        PaymentStatus = "new"
	PaymentAuthorized PaymentStatus = "authorized"
	PaymentPaid       PaymentStatus = "paid"
	PaymentDisbursed  PaymentStatus = "disbursed"
	PaymentCancelled  PaymentStatus = "cancelled"
)

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
