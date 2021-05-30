package clients

type Service struct {
	AddUseCase    AddUseCase
	ListUseCase   ListUseCase
	GetUseCase    GetUseCase
	UpdateUseCase UpdateUseCase
	DeleteUseCase DeleteUseCase
}

func NewService(a AddUseCase, l ListUseCase, g GetUseCase, u UpdateUseCase, d DeleteUseCase) Service {
	return Service{a, l, g, u, d}
}
