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

var countid int = 1 //default emp id is 1

func AddEmp() {
	in := bufio.NewReader(os.Stdin)
	//countid := len(a) + 1
	countid++ // Increment countid based on existing employees

	fmt.Println("Add Emp Details")
	fmt.Println("Enter to start")
	in.ReadString('\n')

	fmt.Println("Note your ID plz id=", countid)

	var firstname string
AddFirstName:
	fmt.Println("Add Emp firstname:")
	input, _ := in.ReadString('\n')
	firstname = strings.TrimSpace(input)
	err := validations.IsValidFirstName(firstname)
	if err != nil {
		fmt.Println(err)
		goto AddFirstName
	}

	var lastname string
AddLastName:
	fmt.Println("Add Emp Lastname:")
	input, _ = in.ReadString('\n')
	lastname = strings.TrimSpace(input)
	err = validations.IsValidLastName(lastname)
	if err != nil {
		fmt.Println(err)
		goto AddLastName
	}

	var email string
Email:
	fmt.Print("Enter email: ")
	input, _ = in.ReadString('\n')
	email = strings.TrimSpace(input)
	err = validations.IsValidEmail(email)
	if err != nil {
		fmt.Println(err)
		goto Email
	}
	err = validations.IsUniqueEmail(email)
	if err != nil {
		fmt.Println(err)
		goto Email
	}

	var password string
Password:
	fmt.Print("Enter Password: ")
	input, _ = in.ReadString('\n')
	password = strings.TrimSpace(input)
	err = validations.IsValidPassword(password)
	if err != nil {
		fmt.Println(err)
		goto Password
	}

	var phoneno string
PhoneNo:
	fmt.Print("Enter PhoneNo: ")
	input, _ = in.ReadString('\n')
	phoneno = strings.TrimSpace(input)
	_, err = validations.IsValidPhoneno(phoneno)
	if err != nil {
		fmt.Println(err)
		goto PhoneNo
	}

	fmt.Println("Add Emp role:")
	role := utility.SelectRole()
	in.ReadString('\n')

	var salary float64
Salary:
	fmt.Print("Enter salary: ")
	input, _ = in.ReadString('\n')
	sal := strings.TrimSpace(input)
	salary, err = validations.IsValidSalary(sal)
	if err != nil {
		fmt.Println(err)
		goto Salary
	}

	var bday time.Time
Bday:
	fmt.Println("Enter the Date in format (YYYY-MM-DD) :")
	dateString, _ := in.ReadString('\n')
	dateString = strings.TrimSpace(dateString) //removing extra charfrom string

	bday, err = validations.IsValidDate(dateString)
	if err != nil {
		fmt.Println(err)
		goto Bday
	}

	// view.A = append(view.A, view.Employee{countid, firstname, lastname, email, password, phoneno, role, salary, bday})
	//err-- main/view.Employee struct literal uses unkeyed fieldscompositesdefault
	//reason -- create instance of view.Employee without using field names
	//sol -- initializestruct using field names
	view.Emp = append(view.Emp, view.Employee{
		ID:        countid,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  password,
		PhoneNo:   phoneno,
		Role:      role,
		Salary:    salary,
		Birthday:  bday,
	})

	fmt.Println("sucessfully enterd emp detais", view.Emp)

}
