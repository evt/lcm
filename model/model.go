package model

// Employee is an organization tree node
type Employee struct {
	Name      string      // employee name
	Employees []*Employee // child employees
}

// AddEmployees adds child employees to a parent employee (=manager)
func (e *Employee) AddEmployees(employees ...*Employee) {
	e.Employees = append(e.Employees, employees...)
}
