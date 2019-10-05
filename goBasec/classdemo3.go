package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person Person) getNameAndAge() (string, int) {
	return person.name, person.age
}

type Student struct {
	Person
	speciality string
}

func (student Student) getSpeciality() string {
	return student.speciality
}
func main() {
	student := new(Student)
	student.name = "zhangsan"
	student.age = 26
	name, age := student.getNameAndAge()
	fmt.Println(name, age)
}
