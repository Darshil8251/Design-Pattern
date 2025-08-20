package organization

type devOps struct {
	name       string
	salary     float64
	position   string
	department string
}

var _ employee = (*devOps)(nil)

func NewDevOps(name string, salary float64, position string, department string) *devOps {
	return &devOps{
		name:       name,
		salary:     salary,
		position:   position,
		department: department,
	}
}

func (d *devOps) GetDetails() string {
	return "Name: " + d.name + ", Position: " + d.position + ", Department: " + d.department
}

func (d *devOps) GetSalary() float64 {
	return d.salary
}

func (d *devOps) GetPosition() string {
	return d.position
}

func (d *devOps) GetDepartment() string {
	return d.department
}
