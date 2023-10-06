package controller

import (
	"fmt"
	"main/validations"
)

// del functinality sathi add kelay this ver same user del own no then to come out
var CurrentuserID int

func Login() {
	fmt.Println() //to get space in cmd
	fmt.Println()
	start := true
	for start {
		fmt.Println("Welcome To Employee Management System")
		fmt.Println("Login")
		emp, ok := validations.AuthUser() //calling AuthUser function cheking enter details are valid or not
		fmt.Println(emp.Role)
		if !ok {
			fmt.Println("Plz Provise Proper Details")
		} else {
			CurrentuserID = emp.ID
			if emp.Role == "Admin" {
				AdminControl() //calling admincontrol func
			} else {
				NonAdminControl(emp) //if role nor admin then calling nonadmin
			}
		}
		fmt.Println()
		fmt.Println()
	}

}
