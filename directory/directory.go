package directory

import "github.com/evt/lcm/model"

// Directory is a corporate directory
type Directory interface {
	GetLowestCommonManager(top, employee1, employee2 *model.Employee) *model.Employee
}
