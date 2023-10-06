package controller

import (
	"bufio"
	"fmt"
	"main/utility"
	"main/validations"
	"main/view"
	"os"
	"strings"
	"time"
)

func UpdateEmp() {

	in := bufio.NewReader(os.Stdin)
	fmt.Println("'Enter' key to begin opration...")
	in.ReadString('\n')
	i, ok := utility.IsEmpPresent()
	if ok {
		EditEmployee(i)
	} else {
		fmt.Println("Employee is Not Present")
	}
}

func EditEmployee(employeeID int) { //editing all values of employee
	countid := employeeID

	fmt.Println("Current Information of Employee")
	e, _ := utility.GetEmployeeByID(employeeID)
	fmt.Printf("ID: %d, Name: %s %s, Email: %s, Role: %s, Salary: %.2f, Birthday: %s\n", e.ID, e.FirstName, e.LastName, e.Email, e.Role, e.Salary, e.Birthday)

	fmt.Println("Edit Emp Details")
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Enter to start")
	in.ReadString('\n')

	var firstname string
EditFirstName:
	fmt.Println("Edit Emp firstname")
	input, _ := in.ReadString('\n')
	firstname = strings.TrimSpace(input)
	err := validations.IsValidFirstName(firstname)
	if err != nil {
		fmt.Println(err)
		goto EditFirstName
	}

	var lastname string
EditLastName:
	fmt.Println("Edit Emp Lastname")
	input, _ = in.ReadString('\n')
	lastname = strings.TrimSpace(input)
	err = validations.IsValidLastName(lastname)
	if err != nil {
		fmt.Println(err)
		goto EditLastName
	}

	var email string
EditEmail:
	fmt.Print("Edit email: ")
	input, _ = in.ReadString('\n')
	email = strings.TrimSpace(input)
	err = validations.IsValidEmail(email)
	if err != nil {
		fmt.Println(err)
		goto EditEmail
	}

	var password string
EditPassword:
	fmt.Print("Edit Password: ")
	input, _ = in.ReadString('\n')
	password = strings.TrimSpace(input)
	err = validations.IsValidPassword(password)
	if err != nil {
		fmt.Println(err)
		goto EditPassword
	}

	var phoneno string
EditPhoneNo:
	fmt.Print("Edit PhoneNo: ")
	input, _ = in.ReadString('\n')
	phoneno = strings.TrimSpace(input)
	err = validations.IsValidPassword(phoneno)
	if err != nil {
		fmt.Println(err)
		goto EditPhoneNo
	}

	fmt.Println("Edit Emp role")
	role := utility.SelectRole()

	var salary float64
EditSalary:
	fmt.Print("Edit salary: ")
	input, _ = in.ReadString('\n')
	sal := strings.TrimSpace(input)
	salary, err = validations.IsValidSalary(sal)
	if err != nil {
		fmt.Println(err)
		goto EditSalary
	}

	var bday time.Time
EditBday:
	fmt.Println("Enter the Date in format (YYYY-MM-DD)")
	dateString, _ := in.ReadString('\n')
	dateString = strings.TrimSpace(dateString)

	bday, err = validations.IsValidDate(dateString)
	if err != nil {
		fmt.Println(err)
		goto EditBday
	}

	updatedEmployee := view.Employee{
		ID:        countid,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  password,
		PhoneNo:   phoneno,
		Role:      role,
		Salary:    salary,
		Birthday:  bday,
	}
	view.Emp[countid-1] = updatedEmployee
	fmt.Printf("Employee with ID %d updated.\n", employeeID)
}
