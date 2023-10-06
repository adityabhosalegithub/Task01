package utility

import "main/view"

func GetEmployeeByID(employeeID int) (view.Employee, bool) {
	//Linear search
	// for _, emp := range a { //itrare karto to check it prestent or not
	// 	if emp.ID == employeeID {
	// 		return emp, true
	// 	}
	// }

	//Binary Search
	left, right := 0, len(view.Emp)-1
	for left <= right {
		mid := left + (right-left)/2

		if view.Emp[mid].ID == employeeID {
			return view.Emp[mid], true
		}

		if view.Emp[mid].ID < employeeID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return view.Employee{}, false
}
