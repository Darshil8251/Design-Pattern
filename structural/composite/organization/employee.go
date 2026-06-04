package organization

type employee interface {
	GetDetails() string
	GetSalary() float64
	GetPosition() string
	GetDepartment() string
}
