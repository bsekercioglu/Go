package structs

import "fmt"

type Person struct {
	firstName     string
	lastName      string
	fatherName    string
	motherName    string
	citizenShipID int64
}

type Student struct {
	identity Person
	class    int
}

func (c Student) save() {
	fmt.Println("Eklendi")
	fmt.Println(c.identity.fatherName)
}

func Demo2() {
	st := Student{class: 1, identity: Person{firstName: "Burak", lastName: "Şekercioğlu", fatherName: "XYZ", motherName: "YXZ", citizenShipID: 12345678901}}
	st.save()
}
