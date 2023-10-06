package main

import "fmt"

type person struct {
	name    string
	phone   int
	address string
}
type employee struct {
	person
	id     int
	dept   string
	salary int
}

func (e *employee) setinfo() {
	fmt.Scan(&e.name, &e.phone, &e.address, &e.id, &e.dept, &e.salary)
}
func (e *employee) getinfo() {
	fmt.Println("name is", e.name)
	fmt.Println("phone no is ", e.phone)
	fmt.Println("address is", e.address)
	fmt.Println("id is", e.id)
	fmt.Println("dept is", e.dept)
	fmt.Println("salary is", e.salary)

}
func main() {
	emp := employee{}
	emp.setinfo()
	emp.getinfo()
}
