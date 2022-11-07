package clients

type Service struct {
	AddUseCase    AddUseCase
	ListUseCase   ListUseCase
	GetUseCase    GetUseCase
	UpdateUseCase UpdateUseCase
	DeleteUseCase DeleteUseCase
}

type Logger interface {
	Info(format string, i ...string)
	Warn(format string, i ...string)
	Err(format string, i ...string)
	Debug(format string, i ...string)
}

func NewService(a AddUseCase, l ListUseCase, g GetUseCase, u UpdateUseCase, d DeleteUseCase) Service {
	return Service{a, l, g, u, d}
}
