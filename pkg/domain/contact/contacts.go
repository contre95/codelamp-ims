package contact

type Repo interface {
	Add(name, phone string) error
	Delete() error
}

type Service interface {
	Create(name, phone string) error
	List() ([]Contact, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo}
}

func (s *service) Create(name, phone string) error {
	if len(name) > 10 {
		panic("name can not be longer than 10")
	}
	err := s.repo.Add(name, phone)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) List() ([]Contact, error) {
	panic("Implement me !")
}
