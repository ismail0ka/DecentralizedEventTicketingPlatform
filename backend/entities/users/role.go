package users

type Role int

const (
	user Role = iota
	organizer
	admin
)

func (r Role) String() string {
	return [...]string{"User", "Organizer", "Admin"}[r]
}