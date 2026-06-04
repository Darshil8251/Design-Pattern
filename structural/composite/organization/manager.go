package organization

import "strings"

type manager struct {
	name       string
	salary     float64
	position   string
	department string
	team       []employee // Managers can handle multiple employees
}

var _ employee = (*manager)(nil)

func NewManager(name string, salary float64, position string, department string) *manager {
	return &manager{
		name:       name,
		salary:     salary,
		position:   position,
		department: department,
		team:       []employee{},
	}
}

func (m *manager) GetDetails() string {
	var details strings.Builder
	details.WriteString("Name: " + m.name + ", Position: " + m.position + ", Department: " + m.department)
	if len(m.team) > 0 {
		details.WriteString(", Team Members: ")
		for _, emp := range m.team {
			details.WriteString(emp.GetDetails() + "; ")
		}
	}
	return details.String()
}

func (m *manager) GetSalary() float64 {
	return m.salary
}

func (m *manager) GetPosition() string {
	return m.position
}

func (m *manager) GetDepartment() string {
	return m.department
}

func (m *manager) AddEmployee(emp employee) {
	m.team = append(m.team, emp)
}
