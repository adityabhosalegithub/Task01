package utility

import (
	"fmt"
	"main/view"
	"time"
)

func UpcomingBirthdays() {
	currentMonth := time.Now().Month()

	var upcoming []view.Employee

	for _, emp := range view.Emp {
		if emp.Birthday.Month() == currentMonth { //		// Check if birthday month matches the current month
			if emp.Birthday.Day() >= time.Now().Day() { // Check if the day is upcoming
				upcoming = append(upcoming, emp)
			}
		}
	}

	if len(upcoming) > 0 {
		for _, alabday := range upcoming {
			fmt.Printf("ID: %d, Name: %s %s, Email: %s, Role: %s, Salary: %.2f\n", alabday.ID, alabday.FirstName, alabday.LastName, alabday.Email, alabday.Role, alabday.Salary)
		}
	} else {
		fmt.Println("No Employee With Upcoming Birthdays")
	}
	// upcoming = upcoming[:0] //making empty
}
