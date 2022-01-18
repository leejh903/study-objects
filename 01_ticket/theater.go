package _01_ticket

type Theater struct {
	ticketSeller TicketSeller
}

func NewTheater(ticketSeller TicketSeller) Theater {
	return Theater{
		ticketSeller: ticketSeller,
	}
}

func (t Theater) Enter(audience Audience) {
	if audience.Bag.HasInvitation() {
		ticket := t.ticketSeller.TicketOffice.GetTicket()
		audience.Bag.Ticket = &ticket
	} else {
		ticket := t.ticketSeller.TicketOffice.GetTicket()
		audience.Bag.MinusAmount(ticket.Fee)
		t.ticketSeller.TicketOffice.PlusAmount(ticket.Fee)
		audience.Bag.Ticket = &ticket
	}
}
