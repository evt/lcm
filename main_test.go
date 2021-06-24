package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite is our test suite
type TestSuite struct {
	suite.Suite
}

// createEmployees creates employee for every letter in alphabet
func createEmployees() map[rune]*Employee {
	org := map[rune]*Employee{}
	for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		org[r] = &Employee{
			name:      string(r),
			employees: []*Employee{},
		}
	}
	return org
}

// Test1 represents the following case:
//
// topManager = Node A
// Employee 1 = Node E
// Employee 2 = Node I
//                A
//             /     \
//            B       C
//          /   \   /   \
//         D     E F     G
//       /   \
//      H     I
//
// Lowest Common Manager for H and E must be B

func (suite *TestSuite) Test1() {
	employees := createEmployees()

	employees['A'].addEmployees(employees['B'], employees['C'])
	employees['B'].addEmployees(employees['D'], employees['E'])
	employees['C'].addEmployees(employees['F'], employees['G'])
	employees['D'].addEmployees(employees['H'], employees['I'])

	lcm := GetLowestCommonManager(employees['A'], employees['E'], employees['H'])

	assert.Equal(suite.T(), lcm, employees['B'])
}

// TestLCM runs our test suite
func TestLCM(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
