package health

// HealthyHandler is in charge of health checking the service
type Service interface {
	Ping() string
}

type service struct {
}

// NewService returns a new HealthHandler
func NewService() Service {
	return &service{}
}

func (s service) Ping() string {
	return "pong"
}
