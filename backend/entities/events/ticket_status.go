package events

type TicketStatus int

const (
	pending TicketStatus = iota
	approved
	rejected
	checkedIn
)

func (ts TicketStatus) String() string {
	return [...]string{"Pending", "Approved", "Rejected", "Checked in"}[ts]
}