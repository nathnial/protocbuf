package main

import "qimiaoPB/pb/person"

func main() {
	var p person.Person
	v := p.HighestEducation.(*person.Person_MiddleSchool)
	v.MiddleSchool = "nono pri school"
}
