package contacts

type Repository interface {
	AddContact(name, phone string) error
	DeleteContact() error
	//Don't worry about the double naming (contact package and contactXxxx methods)
	//because outside this package you will instantiate a svc := NewService()...
	//and then the calls to this methods will be something like this:
	//svc.AddContact("Lucas", "Trolazo")
}

type Service interface {
	Create(name, phone string) error
	List() ([]Contact, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(name, phone string) error {
	if len(name) > 10 {
		panic("name can not be longer than 10")
	}
	err := s.repo.AddContact(name, phone)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) List() ([]Contact, error) {
	panic("Implement me !")
}
