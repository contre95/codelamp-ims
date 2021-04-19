package contacts

type ContactID uint

type Contact struct {
	ID        ContactID
	FirstName string
	LastName  string
	Phone     string
	Comments  string
}
