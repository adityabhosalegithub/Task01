package controller

import (
	"bufio"
	"fmt"
	"main/utility"
	"main/validations"
	"main/view"
	"os"
	"strings"
)

func UpdateEmpField() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("'Enter' key to begin opration...")
	in.ReadString('\n')
	i, ok := utility.IsEmpPresent()
	if ok {
		EditEmployeeField(i)
	} else {
		fmt.Println("Employee is Not Present")
	}
}

// func for takeing selected field from user to edit
func EditEmployeeField(employeeID int) {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Current Information of Employee")
	e, _ := utility.GetEmployeeByID(employeeID)
	fmt.Printf("ID: %d, Name: %s %s, Email: %s, Role: %s, Salary: %.2f, Birthday: %s\n", e.ID, e.FirstName, e.LastName, e.Email, e.Role, e.Salary, e.Birthday)
	in.ReadString('\n')

	fmt.Println("Select Field to Edit:")
	fmt.Println("1. First Name")
	fmt.Println("2. Last Name")
	fmt.Println("3. Email")
	fmt.Println("4. Password")
	fmt.Println("5. Phone Number")
	fmt.Println("6. Role")
	fmt.Println("7. Salary")
	fmt.Println("8. Birthday")

	var choice int
	_, _ = fmt.Scanf("%d", &choice) //fmt.Scanf reads the integer but leaves the newline character in the input buffer

	switch choice {
	case 1:
		EditField(&e.FirstName, "First Name")
	case 2:
		EditField(&e.LastName, "Last Name")
	case 3:
		EditField(&e.Email, "Email")
	case 4:
		EditField(&e.Password, "Password")
	case 5:
		EditField(&e.PhoneNo, "Phone Number")
	case 6:
		e.Role = utility.SelectRole()
	case 7:
		fmt.Print("Enter salary: ")
		input, _ := in.ReadString('\n')
		sal := strings.TrimSpace(input)
		salary, err := validations.IsValidSalary(sal)
		if err != nil {
			fmt.Println(err)
			// goto Salary
		}
		e.Salary = salary
	case 8:
		fmt.Println("Enter the Date in format (YYYY-MM-DD)")
		dateString, _ := in.ReadString('\n')
		dateString = strings.TrimSpace(dateString) //removing extra charfrom string

		bday, err := validations.IsValidDate(dateString)
		if err != nil {
			fmt.Println(err)
		}
		e.Birthday = bday
	case 9:
		break
	default:
		fmt.Println("Invalid choice. Please select a valid option")

	}

	view.Emp[employeeID-1] = e
	fmt.Printf("Employee with ID %d updated.\n", employeeID)
}

// we modify the variable in the calling function,so ue used pointer to it.
// err before use pointer is"argument field is overwritten before first use (SA4009)"
func EditField(field *string, fieldName string) {
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Edit %s: ", fieldName)
	in.ReadString('\n')
	input, _ := in.ReadString('\n')
	newVal := strings.TrimSpace(input)
	*field = newVal //asssigning value to specific field
	fmt.Printf("%s updated to: %s\n", fieldName, *field)
}
