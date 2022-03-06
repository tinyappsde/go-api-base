package models

import (
	"tinyapps/api-base/handlers"
)

type Company struct {
	Id			int			`json:"id"`
	Name		string		`json:"name"`
	Employees	[]Employee	`json:"employees"`
}

func DummyCompany() Company {
	return Company{Name: "Lorem Ipsum Inc.", Employees: []Employee{DummyEmployee()}}
}

func GetCompanies(env *handlers.Env, limit int) []Company {
	results, err := env.DB.Query("SELECT `id`, `name` FROM `companies` LIMIT ?", limit)
	if err != nil {
		panic(err.Error())
	}

	companies := []Company{}

	for results.Next() {
		var company Company
		err = results.Scan(&company.Id, &company.Name)
		if err != nil {
			panic(err.Error())
		}

		company.Employees = GetEmployeesOfCompany(env, company, 10)
		companies = append(companies, company)
	}

	return companies
}

func SaveCompany(env *handlers.Env, company Company) {
	res, err := env.DB.Exec("INSERT INTO `companies` SET `name` = ?", company.Name)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	company.Id = int(id)

	for _, employee := range company.Employees {
		SaveEmployee(env, employee, company)
	}
}
