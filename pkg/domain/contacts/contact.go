package contacts

type ContactID int

type Contact struct {
	ID        ContactID
	FirstName string
	LastName  string
	Phone     string
	Comments  string
}
