package departments

type Developer struct{}
type PM struct{}
type HR struct{}

func (d Developer) GetDepartmentName() string { // implements Department interface
	return "software development"
}

func (pm PM) GetDepartmentName() string {
	return "product management"
}

func (hr HR) GetDepartmentName() string {
	return "human resources"
}
