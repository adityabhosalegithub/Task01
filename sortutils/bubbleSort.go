package sortutils

import (
	"main/view"
	"strings"
)

// implementing bubble sort algo to sort emp by name
func BubbleSortByName(employees []view.Employee) []view.Employee {
	if len(employees) <= 1 {
		return employees
	}
	count := len(employees)
	for i := 0; i < count; i++ {
		flag := false
		for j := 0; j < count-i-1; j++ {
			//in every sucessful loop large element come out like bubble
			if strings.ToLower(employees[j].FirstName) > strings.ToLower(employees[j+1].FirstName) {
				//in if case we checking by converting first name in lower case
				temp := employees[j]
				employees[j] = employees[j+1]
				employees[j+1] = temp
				flag = false
			}

			if flag {
				break
			}
		}
	}
	return employees
}

// implementing bubble sort algo to sort emp by role
func BubbleSortByRole(employees []view.Employee) []view.Employee {
	if len(employees) <= 1 {
		return employees
	}
	count := len(employees)
	for i := 0; i < count; i++ {
		flag := false
		for j := 0; j < count-i-1; j++ {
			if employees[j].Role > employees[j+1].Role {
				temp := employees[j]
				employees[j] = employees[j+1]
				employees[j+1] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return employees
}
