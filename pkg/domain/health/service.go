package health

// HealthyHandler is in charge of health checking the service
type Service interface {
	Ping() string
}

type Logger interface {
	Info(format string, i ...interface{})
	Warn(format string, i ...interface{})
	Err(format string, i ...interface{})
	Debug(format string, i ...interface{})
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
