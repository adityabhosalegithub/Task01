/*
	View his/her details

- Update his/her details
- Search employee(display firstname, lastname, email, role)
*/
package controller

import (
	"fmt"
	"main/sortutils"
	"main/validations"
	"main/view"
)

func NonAdminControl(e view.Employee) {
	var id int = e.ID
	start := true
	for start {
		fmt.Println() //to get space in cmd
		fmt.Println()
		fmt.Println("SELECT OPTION")
		fmt.Println("1.Update Own All Data") //GetEmployeeByID(employeeID int) (Employee, bool)
		fmt.Println("2.Update Own Selected Fields")
		fmt.Println("3.See AllEmp Data ")
		fmt.Println("4.LogOut")

		var i int
		_, err := fmt.Scanf("%d", &i)
		validations.CheckNilErr(err)

		switch key := i; key {
		case 1:
			EditEmployee(id) //passing id
		case 2:
			EditEmployeeField(id)
		case 3:
			fmt.Println("Loding all Emp.....")
			sortutils.NonAdminPrintAllEmp()
		case 4:
			fmt.Println("Sucessful LogOut .....")
			start = false
		default:
			fmt.Println("Invalid plze select Proper Option")
		}
		fmt.Println()
	}
}
