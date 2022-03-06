package models

import (
	"tinyapps/api-base/handlers"
)

type Employee struct {
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
}

func DummyEmployee() Employee {
	return Employee{FirstName: "John", LastName: "Doe"}
}

func GetEmployeesOfCompany(env *handlers.Env, company Company, limit int) []Employee {
	results, err := env.DB.Query("SELECT `first_name`, `last_name` FROM `employees` WHERE `company_id` = ? LIMIT ?", company.Id, limit)
	if err != nil {
		panic(err.Error())
	}

	employees := []Employee{}

	for results.Next() {
		var employee Employee
		err = results.Scan(&employee.FirstName, &employee.LastName)
		if err != nil {
			panic(err.Error())
		}

		employees = append(employees, employee)
	}

	return employees
}

func SaveEmployee(env *handlers.Env, employee Employee, company Company) {
	_, err := env.DB.Exec("INSERT INTO `employees` SET `company_id` = ?, `first_name` = ?, `last_name` = ?", company.Id, employee.FirstName, employee.LastName)
	if err != nil {
		panic(err)
	}
}
