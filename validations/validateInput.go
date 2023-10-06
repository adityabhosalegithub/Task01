package validations

import (
	"fmt"
	"main/view"
	"strconv"
	"strings"
	"time"
)

// all cutom validations
func IsValidFirstName(firstname string) error {
	if firstname != "" {
		return nil
	}
	return CustomError{message: "Firstname cannot be empty. Enter Name again"}
	// createing CustomError instance with a specific error message and return
}

func IsValidLastName(lastname string) error {
	if lastname != "" {
		return nil
	}
	return CustomError{message: "Lastname cannot be empty. Enter Name again"}
}

func IsValidEmail(email string) error {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return nil
	}
	return CustomError{message: "Email should contain @ and . Enter valid email"}
}

func IsValidPassword(password string) error {
	if password != "" {
		return nil
	}
	return CustomError{message: "Password cannot be empty. Enter password again"}
}

func IsValidPhoneno(phoneno string) (int, error) {
	if len(phoneno) != 10 {
		return 0, CustomError{message: "phone number must be 10 digits long"}
	}
	phone, err := strconv.ParseInt(phoneno, 10, 64)
	if err != nil {
		return 0, CustomError{message: "failed to convert phone number to integer"}
	}
	return int(phone), nil
}

func IsValidSalary(salary string) (float64, error) {
	if salary == "" {
		return 0, CustomError{message: "Salary cannot be empty. Enter salary again"}
	}
	convertedSalary, err := strconv.ParseFloat(salary, 64)
	if err != nil {
		return 0, CustomError{message: "failed to convert salary to float64"}
	}
	return convertedSalary, nil
}

func IsValidDate(dateString string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Time{}, CustomError{message: "Enterd Date not in proper format .Enter date again"}
	}
	fmt.Printf("You entered: %s\n", date.Format("January 2, 2006"))
	//for checking emp above 18 year
	currentTime := time.Now()
	age := currentTime.Year() - date.Year()
	if currentTime.Before(time.Date(currentTime.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)) {
		age-- // to handle cases where the birthday not occurred yet in the current year.
	}
	if age >= 18 {
		fmt.Println("Employee age is 18 or older.")
	} else {
		return date, CustomError{message: "Employee is not 18 years old yet. Please enter a valid birthdate."}
	}
	return date, nil
}

func CheckNilErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func IsValidInt(input string) (int, error) {
	num, err := strconv.ParseInt(input, 10, 64) //converting string to int
	if err != nil {
		return 0, CustomError{message: "failed to convert phone number to integer"}
	}
	return int(num), nil
}

func IsEmptyString(input string) error {
	str := strings.TrimSpace(input)
	if str == "" {
		return CustomError{message: "Empty value entered...Enter value again"}
	}
	return nil
}

func IsUniqueEmail(email string) error { //to check email is unique
	for _, emp := range view.Emp {
		if emp.Email == email {
			return CustomError{message: "Email AllReady Present"}
		}
	}
	return nil
}

func IsEmptyInt(input int) error {
	if input == 0 {
		return CustomError{message: "Empty value entered...Enter value again"}
	}
	return nil
}
