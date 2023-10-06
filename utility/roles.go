package utility

import (
	"fmt"
	"main/validations"
)

func SelectRole() string {
	start := true
	for start {
		fmt.Println("SELECT ROLE")
		fmt.Println("1.Admin")
		fmt.Println("2.Manager")
		fmt.Println("3.Developer")
		fmt.Println("4.Tester")

		var i int
		_, err := fmt.Scanf("%d", &i)
		validations.CheckNilErr(err)

		switch key := i; key {
		case 1:
			return "Admin"
		case 2:
			return "Manager"
		case 3:
			return "Developer"
		case 4:
			return "Tester"
		default:
			fmt.Println("Invalid plze select Proper Option")
		}
	}
	return ""
}
