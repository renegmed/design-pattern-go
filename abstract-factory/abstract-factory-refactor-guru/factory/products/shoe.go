package products

type shoe struct {
	logo string
	size int
}

func (s *shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *shoe) GetLogo() string {
	return s.logo
}

func (s *shoe) SetSize(size int) {
	s.size = size
}

func (s *shoe) GetSize() int {
	return s.size
}
