package directory

import "github.com/evt/lcm/model"

// DefaultDirectory is a default implementation for Directory interface
type DefaultDirectory struct{}

// Make sure DefaultDirectory implements Directory interface
var _ Directory = (*DefaultDirectory)(nil)

func NewDefaultDirectory() DefaultDirectory {
	return DefaultDirectory{}
}

// GetLowestCommonManager returns lowest common manager for org tree boss (=top), employee1 and employee2
func (d DefaultDirectory) GetLowestCommonManager(top, employee1, employee2 *model.Employee) *model.Employee {
	return getLowestCommonManager(top, employee1, employee2).lowestCommonManager
}

// lcm is just a wrapper around lowest common manager to use in recursive function below
type lcm struct {
	lowestCommonManager *model.Employee
	foundEmployees      int
}

// getLowestCommonManager is a recursive function that does the dirty job :)
func getLowestCommonManager(manager, employee1, employee2 *model.Employee) lcm {
	foundEmployees := 0

	for _, child := range manager.Employees {
		childLCM := getLowestCommonManager(child, employee1, employee2)
		if childLCM.lowestCommonManager != nil {
			return childLCM
		}
		foundEmployees += childLCM.foundEmployees
	}

	if manager == employee1 || manager == employee2 {
		foundEmployees++
	}

	var lcmEmployee *model.Employee

	if foundEmployees == 2 {
		lcmEmployee = manager
	}

	return lcm{
		lowestCommonManager: lcmEmployee,
		foundEmployees:      foundEmployees,
	}
}
