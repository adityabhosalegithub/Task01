package sortutils

import (
	"fmt"
	"main/view"
)

func PrintAllEmp() { // to print all filed for admin role
	for _, emp := range view.Emp {
		fmt.Printf("ID: %d, Full_Name: %s %s, Email: %s, Role: %s, Salary: %.2f ,Birthday: %s\n", emp.ID, emp.FirstName, emp.LastName, emp.Email, emp.Role, emp.Salary, emp.Birthday.Format("January 2, 2006"))
	}
}

// to print specific filed for non admin role
func NonAdminPrintAllEmp() {
	for _, emp := range view.Emp {
		fmt.Printf("ID: %d, Full_Name: %s %s, Email: %s, Role: %s\n", emp.ID, emp.FirstName, emp.LastName, emp.Email, emp.Role)
	}
}
