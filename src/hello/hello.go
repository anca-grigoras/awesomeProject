package main

import (
	"awesomeProject/src/websockets"
	"fmt"
	"log"
	"net/http"
	_ "net/http"
)

func main() {
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)
	//
	//names := []string{"Gladys", "Samantha", "Darrin"}
	//
	//messages, err := greetings.Hellos(names)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(messages)
	//
	//name := letters.MyString("Sam Anderson")
	//var v letters.VowelsFinder
	//v = name
	//fmt.Printf("Vowels are %c\n", v.FindVowels())
	//
	//pemp1 := salary.Permanent{
	//	EmpId:    1,
	//	BasicPay: 5000,
	//	Pf:       20,
	//}
	//pemp2 := salary.Permanent{
	//	EmpId:    2,
	//	BasicPay: 6000,
	//	Pf:       30,
	//}
	//cemp1 := salary.Contract{
	//	EmpId:    3,
	//	BasicPay: 3000,
	//}
	//
	//freelancer1 := salary.Freelancer{
	//	EmpId:       4,
	//	RatePerHour: 70,
	//	TotalHours:  120,
	//}
	//freelancer2 := salary.Freelancer{
	//	EmpId:       5,
	//	RatePerHour: 100,
	//	TotalHours:  100,
	//}
	//
	//employees := []salary.SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
	//salary.TotalExpense(employees)

	fmt.Println("Hello World")
	websockets.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
