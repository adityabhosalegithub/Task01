package controller

import (
	"bufio"
	"fmt"
	"main/sortutils"
	"main/utility"
	"main/validations"
	"os"
)

func AdminControl() {
	in := bufio.NewReader(os.Stdin)
	start := true
	for start {
		fmt.Println() //to get space in cmd
		fmt.Println()
		fmt.Println("###...ADMIN PAGE...###")
		fmt.Println("Select Options")
		fmt.Println("'Enter' key to begin opration...")
		in.ReadString('\n')

		fmt.Println("1. Add Employee")
		fmt.Println("2. Delete Employee")
		fmt.Println("3. Update All Data OfEmployee")
		fmt.Println("4. Update Employee Data with Specific Fields")
		fmt.Println("5. Check if Employee is Present")
		fmt.Println("6. Get Employee by ID")
		fmt.Println("7. Print All Employees")
		fmt.Println("8. Employees with Upcoming Birthday")
		fmt.Println("9. Print Employees According To Specific Fields")
		fmt.Println("10. Log Out")

		var i int
		_, err := fmt.Scanf("%d", &i)
		validations.CheckNilErr(err)

		switch key := i; key {
		case 1:
			AddEmp()
		case 2:
			DelEmp()
		case 3:
			UpdateEmp()
		case 4:
			UpdateEmpField()
		case 5:
			fmt.Println("'Enter' key to begin opration...")
			in.ReadString('\n')

			fmt.Println("Enter Emp No ")
			var j int
			_, err := fmt.Scanf("%d", &j)
			validations.CheckNilErr(err)
			_, ok := utility.GetEmployeeByID(j) //return bool  //internammly getempbyid
			if ok {
				fmt.Println("Emp is Present")
			} else {
				fmt.Println("No match found")
			}
		case 6:

			fmt.Println("'Enter' key to begin opration...")
			fmt.Println("Enter Emp No:")
			in.ReadString('\n')

			var j int
			_, err := fmt.Scanf("%d", &j)
			validations.CheckNilErr(err)
			e, ok := utility.GetEmployeeByID(j) //retrun emp ,bool
			if ok {
				fmt.Printf("ID: %d, Name: %s %s, Email: %s, Role: %s, Salary: %.2f\n", e.ID, e.FirstName, e.LastName, e.Email, e.Role, e.Salary)
			} else {
				fmt.Println("No match found")
			}

		case 7:
			fmt.Println("Loding all Emp.....")
			sortutils.PrintAllEmp()
		case 8:
			fmt.Println("List(Print) employees who have upcoming birthday in this month")
			utility.UpcomingBirthdays()
		case 9:
			fmt.Println("Loding all Emp.....")
			sortutils.PrintSorted()
		case 10:
			fmt.Println("Sucessful LogOut .....")
			start = false
		default:
			fmt.Println("Invalid plze select Proper Option")
		}
		fmt.Println()
		fmt.Println()
	}

}
