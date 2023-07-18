package services

type Services struct {
	*PotatoService
}

func NewServices() *Services {
	return &Services{
		PotatoService: NewPotatoService(),
	}
}
