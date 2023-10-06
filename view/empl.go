package view

import "time"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	PhoneNo   string
	Role      string
	Salary    float64
	Birthday  time.Time
}

var Emp []Employee // a global var of type slice of Employee

//adding default data of Admin
var DefaultEmp Employee

func init() {
	DefaultEmp = Employee{
		ID:        1,
		FirstName: "TakeOff",
		LastName:  "Technology",
		Email:     "admin@techoff.com",
		Password:  "1111",
		PhoneNo:   "1234567890",
		Role:      "Admin",
		Salary:    50000.0,
		Birthday:  time.Date(1990, 10, 20, 0, 0, 0, 0, time.UTC),
	}

	Emp = append(Emp, DefaultEmp)
}
