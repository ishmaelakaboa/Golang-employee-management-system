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

func (e *Employee) UpdateSalary(newSalary float64){
	if newSalary > 0 {
		e.Salary = newSalary
		fmt.Printf("Updated salary for %s to $%.2f\n", e.Name, e.Salary)
	} else {
		fmt.Println("Invalid salary amount.")
	}
}

func (e *Employee) Promote(newPosition string, raise float64){
	e.Position = newPosition
	e.UpdateSalary(e.Salary + raise)
	fmt.Printf("%s has been promoted to %s!\n", e.Name, e.Position)
}

type EmployeeManagementSystem struct {
	Employees []Employee
}

func (ems *EmployeeManagementSystem) AddEmployee(employee Employee) {
	ems.Employees = append(ems.Employees, employee)
	fmt.Println("Added new employee:", employee.Name)
}

func (ems EmployeeManagementSystem) FindEmployeeByID(id int) *Employee{
	for i, emp := range ems.Employees{
		if emp.ID == id {
			return &ems.Employees[i]
		}
	}
	return nil
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
		emp.UpdateSalary(65000)
		emp.Promote("Senior Software Engineer", 10000)
	} else {
		fmt.Println("Employee not found.")
	}

	//Display updated employee list
	fmt.Println("\nAfter updates:")
	ems.ListEmployees()
}