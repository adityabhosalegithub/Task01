package validations

import (
	"bufio"
	"fmt"
	"main/view"
	"os"
	"strings"
)

func AuthUser() (view.Employee, bool) { //func to take input (id,Pass) from user and validate
	in := bufio.NewReader(os.Stdin)

ID:
	fmt.Println("Enter Employee_ID :")
	str, _ := in.ReadString('\n')
	str = strings.TrimSpace(str)
	if str == "" {
		fmt.Println("Input is empty. Please enter a number.")
		goto ID
	}
	empid, err := IsValidInt(str)
	if err != nil {
		fmt.Println(err)
	}

Pass:
	fmt.Print("Enter Password: ")
	input, _ := in.ReadString('\n')
	password := strings.TrimSpace(input)
	err = IsValidPassword(password)
	if err != nil {
		fmt.Println(err)
		goto Pass
	}

	for _, emp := range view.Emp { //itrare to check emp present or not
		if emp.ID == empid {
			if emp.Password == password {
				return emp, true
			}
		}
	}
	return view.Employee{}, false
}
