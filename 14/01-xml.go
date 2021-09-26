package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName     xml.Name `xml:"person"`
	Id          int      `xml:"id,attr"`
	Name        string   `xml:"name"`
	Gender      string   `xml:"gender"`
	Nationality string   `xml:"nationality"`
	Resident    []string `xml:"resident"`
}

func (p Person) String() string {
	return fmt.Sprintf("Person id=%v, name=%v, gender=%v, nationality=%v", p.Id, p.Name, p.Gender, p.Nationality)
}

func main() {
	person1 := &Person{
		Id:          1,
		Name:        "Kunal",
		Gender:      "male",
		Nationality: "India",
		Resident:    []string{"Bihar", "Haryana", "Maharashtra", "Uttar Pradesh", "West Bengal"},
	}

	output, err := xml.MarshalIndent(person1, " ", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
	// Print with XML Header
	fmt.Println(xml.Header + string(output))

	// Unmarshaling xml data
	var uPerson1 Person
	err = xml.Unmarshal(output, &uPerson1)
	if err != nil {
		panic(err)
	}
	fmt.Println(uPerson1)

}
