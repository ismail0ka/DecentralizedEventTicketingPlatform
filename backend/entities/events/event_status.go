package events

type EventStatus int

const (
	available EventStatus = iota
	sold_out
	cancelled
)

func (es EventStatus) String() string {
	return [...]string{"Available", "Sold out", "Cancelled"}[es]
}