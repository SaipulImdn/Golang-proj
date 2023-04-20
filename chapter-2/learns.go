package main

import (
	"fmt"
	"strings"
)

var numberA int = 4
var numberB *int = &numberA

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"syaiful", "rozy", "abdul"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [rizky]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]

	{
		fmt.Println("numberA (value)  :", numberA)   // 4
		fmt.Println("numberA (address) :", &numberA) //0xc20800a220

		fmt.Println("numberB (value) :", *numberB)  //4
		fmt.Println("numberB (address) :", numberB) //0xc20800a220
	}

	{
		fmt.Println("numberA (value)   :", numberA)
		fmt.Println("numberA (address) :", &numberA)
		fmt.Println("numberB (value)   :", *numberB)
		fmt.Println("numberB (address) :", numberB)

		fmt.Println("")

		numberA = 5

		fmt.Println("numberA (value)   :", numberA)
		fmt.Println("numberA (address) :", &numberA)
		fmt.Println("numberB (value)   :", *numberB)
		fmt.Println("numberB (address) :", numberB)
	}

	{
		var s1 sekolah
		s1.name = "syaiful imanudin"
		s1.sekolah = "MAN 2 Kota Cirebon"
		s1.kelas = 2

		fmt.Println("name :", s1.name)
		fmt.Println("sekolah :", s1.sekolah)
		fmt.Println("kelas :", s1.kelas)
	}

	{
		var s1 = student{name: "syaiful", grade: 2}

		var s2 *student = &s1
		fmt.Println("student 1, name :", s1.name)
		fmt.Println("student 4, name :", s2.name)

		s2.name = "syaiful"
		fmt.Println("student 1, name :", s1.name)
		fmt.Println("student 4, name :", s2.name)
	}

	{
		var s1 = student{"syaiful imanudin", 20}
		s1.sayHello()

		var name = s1.getNameAt(2)
		fmt.Println("nama panggilan :", name)
	}

	{
		var s1 = student{"syaiful imanuidn", 21}
		fmt.Println("s1 before", s1.name)
		// john wick

		s1.changeName1("rizky abdullah")
		fmt.Println("s1 after changeName1", s1.name)
		// john wick

		s1.changeName2("kiki abdul")
		fmt.Println("s1 after changeName2", s1.name)
		// ethan hunt
	}
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

type student struct {
	name  string
	grade int
}

type sekolah struct {
	name    string
	sekolah string
	kelas   int
}
