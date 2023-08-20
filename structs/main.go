package main

import "log"

type person struct {
	name string
	age  int32
}

type contactInfo struct {
	zip uint64
}

type newPerson struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p1 := person{"Bhavik", 12}
	p1.print()

	p1.name = "Bhavik Brahambhatt"
	p1.age = 23

	var p2 person
	p2.print()

	p3 := newPerson{
		firstName: "B1", lastName: "B2", contactInfo: contactInfo{zip: uint64(1)},
	}
	p3.print()

	p3.updateName("BB")
	p3.print()
}

func (p person) print() {
	log.Printf("%+v", p)
}

func (p newPerson) print() {
	log.Printf("%+v", p)
}

func (p *newPerson) updateName(name string) {
	(*p).firstName = name
	p2 := *p

	p2.firstName = "BBBB"
	p.print()
	p2.print()
}
