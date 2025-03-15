package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func (e Employee) Display(){
	fmt.Printf("ID: %d\nName: %s\nPosition: %s\nSalary: $%.2f\n", e.ID, e.Name, e.Position, e.Salary)
}

type EmployeeManagementSystem struct {
	Employees []Employee
}

func (ems *EmployeeManagementSystem) AddEmployee(employee Employee) {
	ems.Employees = append(ems.Employees, employee)
	fmt.Println("Added new employee:", employee.Name)
}

func (ems EmployeeManagementSystem) ListEmployees(){
	fmt.Println("Employee List:")
	for _, emp := range ems.Employees{
		emp.Display()
		fmt.Println("__________________________")
	}
}

func main() {
	//initialize system
	ems := EmployeeManagementSystem{}

	//add employees
	ems.AddEmployee(Employee{ID: 1, Name: "Mark Smith", Position: "Sotware Engineer", Salary: 60000})
	ems.AddEmployee(Employee{ID: 2, Name: "Jane White", Position: "Project Manager", Salary: 80000})

	//Display employees
	ems.ListEmployees()

	//find employee and update
	emp := ems.FindEmployeeByID(1)
	if emp != nil {

	} else
}