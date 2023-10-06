package sortutils

import (
	"fmt"
	"main/view"
)

func PrintSorted() {
	//swich case for selecting soting method
	var sortedEmployees []view.Employee
	for {
		fmt.Println("SORT BY....")
		fmt.Println("1.ID")
		fmt.Println("2.NAME")
		fmt.Println("3.ROLE")
		fmt.Println("4.SALARY")
		fmt.Println("5.EXIT")

		var choice int
		fmt.Scan(&choice)
		fmt.Println("selected ", choice)
		if choice == 5 {
			break
		}

		switch choice {
		case 1:
			sortedEmployees = view.Emp //we added emp according to id in increment
		case 2:
			sortedEmployees = BubbleSortByName(view.Emp)
		case 3:
			sortedEmployees = BubbleSortByRole(view.Emp)
		case 4:
			sortedEmployees = SelectionSortBySal(view.Emp)
		default:
			fmt.Println("Invalid plze select Proper Option")
		}

		for _, emp := range sortedEmployees {
			fmt.Printf("ID: %d, Full_Name: %s %s, Email: %s, Role: %s, Salary: %.2f ,Birthday: %s\n", emp.ID, emp.FirstName, emp.LastName, emp.Email, emp.Role, emp.Salary, emp.Birthday.Format("January 2, 2006"))
		}
		sortedEmployees = sortedEmployees[:0]
	}
}
