package _01_ticket

type TicketOffice struct {
	amount  int64
	tickets []Ticket
}

func NewTicketOffice(amount int64, tickets ...Ticket) *TicketOffice {
	return &TicketOffice{
		amount:  amount,
		tickets: tickets,
	}
}

func (o *TicketOffice) GetTicket() Ticket {
	ticket := o.tickets[0]
	o.tickets = o.tickets[1:]
	return ticket
}

func (o *TicketOffice) MinusAmount(amount int64) {
	o.amount -= amount
}

func (o *TicketOffice) PlusAmount(amount int64) {
	o.amount += amount
}
