package main

import (
	"fmt"
	"log"
)

// Employee is an orgaization tree node
type Employee struct {
	name      string      // employee name
	employees []*Employee // child employees
}

// addEmployees adds child employees to a parent employee (=manager)
func (e *Employee) addEmployees(employees ...*Employee) {
	e.employees = append(e.employees, employees...)
}

// LCM is just a wrapper around lowest common manager to use in recursive function below
type LCM struct {
	lowestCommonManager *Employee
	foundEmployees      int
}

// GetLowestCommonManager returns lowest common manager for org tree boss (=top), employee1 and employee2
func GetLowestCommonManager(top, employee1, employee2 *Employee) *Employee {
	return getLowestCommonManager(top, employee1, employee2).lowestCommonManager
}

func getLowestCommonManager(manager, employee1, employee2 *Employee) LCM {
	foundEmployees := 0

	for _, child := range manager.employees {
		childLCM := getLowestCommonManager(child, employee1, employee2)
		if childLCM.lowestCommonManager != nil {
			return childLCM
		}
		foundEmployees += childLCM.foundEmployees
	}

	if manager == employee1 || manager == employee2 {
		foundEmployees++
	}

	var lcmEmployee *Employee

	if foundEmployees == 2 {
		lcmEmployee = manager
	}

	return LCM{
		lowestCommonManager: lcmEmployee,
		foundEmployees:      foundEmployees,
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Printf("Please run `make test` instead ;)")
	return nil
}
