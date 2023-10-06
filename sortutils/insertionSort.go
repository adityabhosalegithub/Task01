package sortutils

import (
	"main/view"
	"strings"
)

// implementing Insertion sort algo to sort emp by role
func InsertionSortByRole(employees []view.Employee) []view.Employee {
	count := len(employees)
	for i := 1; i < count; i++ {
		j := i - 1
		temp := employees[i]
		for j >= 0 && strings.Compare(employees[j].Role, temp.Role) > 0 {
			employees[j+1] = employees[j]
			j--
		}
		employees[j+1] = temp
	}
	return employees
}
