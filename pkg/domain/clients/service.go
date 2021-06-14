package clients

type Service struct {
	AddUseCase    AddUseCase
	ListUseCase   ListUseCase
	GetUseCase    GetUseCase
	UpdateUseCase UpdateUseCase
	DeleteUseCase DeleteUseCase
}

type Logger interface {
	Info(format string, i ...interface{})
	Warn(format string, i ...interface{})
	Err(format string, i ...interface{})
	Debug(format string, i ...interface{})
}

func NewService(a AddUseCase, l ListUseCase, g GetUseCase, u UpdateUseCase, d DeleteUseCase) Service {
	return Service{a, l, g, u, d}
}
