package controller

import (
	"bufio"
	"fmt"
	"main/utility"
	"main/validations"
	"main/view"
	"os"
)

func DelEmp() {
	fmt.Println("Enter Emp No to Be Delete:")
	in := bufio.NewReader(os.Stdin)
	in.ReadString('\n')
	var i int
	_, err := fmt.Scanf("%d", &i)
	validations.CheckNilErr(err)
	_, ok := utility.GetEmployeeByID(i)
	if !ok {
		fmt.Println("please prvide correct information")
	} else {
		DeleteEmployee(i)
		if CurrentuserID == i { //checking current emp id with id which is deleted
			fmt.Println("you Deleted Your Acccount")
			Login() //agin to login function
		}
	}

}
func DeleteEmployee(employeeID int) {
	for i, emp := range view.Emp {
		if emp.ID == employeeID {
			view.Emp = append(view.Emp[:i], view.Emp[i+1:]...) //it come till i-1 and exclude i
			//i+1is start from there it will get append next to that
			fmt.Printf("Employee with ID %d deleted.\n", employeeID)
			return
		}
	}
	fmt.Printf("Employee with ID %d not found.\n", employeeID)
}
