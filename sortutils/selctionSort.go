package sortutils

import "main/view"

// implementing selection sort algo to sort emp by name
func SelectionSortBySal(employees []view.Employee) []view.Employee {
	count := len(employees)
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if employees[i].Salary > employees[j].Salary { //small element first coome out
				temp := employees[j]
				employees[j] = employees[i]
				employees[i] = temp
			}
		}
	}
	return employees
}
