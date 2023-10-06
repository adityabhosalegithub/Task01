package utility

import (
	"fmt"
	"main/validations"
)

func IsEmpPresent() (int, bool) {
No:
	fmt.Println("Enter Emp No. ")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
		goto No
	} else {
		err = validations.IsEmptyInt(i)
		if err != nil {
			fmt.Println(err)
			goto No
		} else {
			//Scanf requires a pointer to the variable where it will store the scanned value.
			_, found := GetEmployeeByID(i)
			if found { //checking emp no is valid or not
				// fmt.Println("Present")
				return i, true
			}
		}
	}
	// fmt.Println("please prvide correct information")
	return i, false
}
