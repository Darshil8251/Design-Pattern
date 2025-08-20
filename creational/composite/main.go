package main

import (
	"design-pattern/creational/composite/organization"
)

func main() {

	// Organization can have developers
	developer := organization.NewDeveloper("Alice", 70000, "Software Engineer", "Engineering")

	// Organization can have DevOps engineers
	devOps := organization.NewDevOps("Bob", 80000, "DevOps Engineer", "Operations")

	// Print details of the employees
	println(developer.GetDetails())
	println(devOps.GetDetails())

	// create a manager who can manage multiple employees
	manager := organization.NewManager("Charlie", 90000, "Engineering Manager", "Engineering")
	manager.AddEmployee(developer)
	manager.AddEmployee(devOps)

	// Print details of the manager and their team
	println(manager.GetDetails())
}
