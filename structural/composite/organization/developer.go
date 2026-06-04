package organization

type developer struct {
	name       string
	salary     float64
	position   string
	department string
}

var _ employee = (*developer)(nil)

func NewDeveloper(name string, salary float64, position string, department string) *developer {
	return &developer{
		name:       name,
		salary:     salary,
		position:   position,
		department: department,
	}
}

func (d *developer) GetDetails() string {
	return "Name: " + d.name + ", Position: " + d.position + ", Department: " + d.department
}

func (d *developer) GetSalary() float64 {
	return d.salary
}

func (d *developer) GetPosition() string {
	return d.position
}

func (d *developer) GetDepartment() string {
	return d.department
}
