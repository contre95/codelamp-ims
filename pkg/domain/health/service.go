package health

// HealthyHandler is in charge of health checking the service
type Service interface {
	Ping() string
}

type Logger interface {
	Info(format string, i ...string)
	Warn(format string, i ...string)
	Err(format string, i ...string)
	Debug(format string, i ...string)
}

type service struct {
	logger Logger
}

// NewService returns a new HealthHandler
func NewService(l Logger) Service {
	return &service{logger: l}
}

func (s service) Ping() string {
	s.logger.Info("Pong")
	return "pong"
}
