package bridge

type Gender interface {
	DescribePerson() string
}

type Department interface {
	GetDepartmentName() string
}
