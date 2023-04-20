package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "syaiful",
	height:      182.5,
	age:         20,
	isGraduated: false,
	hobbies:     []string{"main bola", "tennis"},
}

func main() {
	fmt.Printf("%b\n", data.age)
	// 11010
	fmt.Printf("%c\n", 1400)
	// n
	fmt.Printf("%c\n", 1235)
	// a
	fmt.Printf("%d\n", data.age)
	// 26
	fmt.Printf("%e\n", data.height)
	// 1.825000e+02
	fmt.Printf("%E\n", data.height)
	// 1.825000E+20
	fmt.Printf("%f\n", data.height)
	// 182.500000
	fmt.Printf("%.9f\n", data.height)
	// 182.500000000
	fmt.Printf("%.2f\n", data.height)
	// 182.50
	fmt.Printf("%.f\n", data.height)
	// 182
	fmt.Printf("%e\n", 0.123123123123)
	// 1.231231e-01
	fmt.Printf("%f\n", 0.123123123123)
	// 0.123123
	fmt.Printf("%g\n", 0.123123123123)
	// 0.123123123123
	fmt.Printf("%g\n", 0.12)
	// 0.12
	fmt.Printf("%.5g\n", 0.12)
	// 0.12
	fmt.Printf("%o\n", data.age)
	// 32
	fmt.Printf("%p\n", &data.age)
	// 0x2081be0c0
	fmt.Printf("%q\n", `"name \ height "`)
	// "\" name \\ height \""
	fmt.Printf("%s\n", data.name)
	// syaiful
	fmt.Printf("%t\n", data.isGraduated)
	// false
	fmt.Printf("%T\n", data.name)
	// string
	fmt.Printf("%T\n", data.height)
	// float64
	fmt.Printf("%T\n", data.age)
	// int32
	fmt.Printf("%T\n", data.isGraduated)
	// bool
	fmt.Printf("%T\n", data.hobbies)
	// []string
	fmt.Printf("%v\n", data)
	// {Syaiful 182.5 21 false [mamin bola tennis]}
	fmt.Printf("%+v\n", data)
	fmt.Printf("%#v\n", data)
	fmt.Printf("%x\n", data.age)

	var data = struct {
		name   string
		height float64
	}{
		name:   "syaiful",
		height: 182.5,
	}
	fmt.Printf("%#v\n", data)

	var d = data.name

	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
	fmt.Printf("%x\n", d)
	fmt.Printf("%%\n")

}
