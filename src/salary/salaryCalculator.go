package salary

import (
	"fmt"
)

type SalaryCalculator interface {
	calculateSalary() int
}

type Permanent struct {
	EmpId    int
	BasicPay int
	Pf       int
}

type Contract struct {
	EmpId    int
	BasicPay int
}

type Freelancer struct {
	EmpId       int
	RatePerHour int
	TotalHours  int
}

// CalculateSalary salary of permanent employee is the sum of basic pay and pf
func (p Permanent) calculateSalary() int {
	return p.BasicPay + p.Pf
}

// CalculateSalary salary of contract employee is the basic pay alone
func (c Contract) calculateSalary() int {
	return c.BasicPay
}

// CalculateSalary salary of freelancer
func (f Freelancer) calculateSalary() int {
	return f.RatePerHour * f.TotalHours
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func TotalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}
