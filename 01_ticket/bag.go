package _01_ticket

type Bag struct {
	amount     int64
	invitation *Invitation
	Ticket     *Ticket
}

func NewBag(amount int64) *Bag {
	return &Bag{
		amount: amount,
	}
}

func NewBagWithInvitation(amount int64, invitation Invitation) *Bag {
	return &Bag{
		amount:     amount,
		invitation: &invitation,
	}
}

func (b Bag) HasInvitation() bool {
	return b.invitation != nil
}

func (b Bag) HasTicket() bool {
	return b.Ticket != nil
}

func (b *Bag) MinusAmount(amount int64) {
	b.amount -= amount
}

func (b *Bag) PlusAmount(amount int64) {
	b.amount += amount
}
