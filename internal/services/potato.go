package services

type Potato struct {
	Name string `json:"name"`
}

type PotatoService struct{}

func (s *PotatoService) GetPotato(name string) Potato {
	return Potato{Name: name}
}

func NewPotatoService() *PotatoService {
	return &PotatoService{}
}
