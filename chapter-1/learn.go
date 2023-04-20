package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var decimalNumber = 2.62
var exist bool = true
var value = (((2+6)%3)*4 - 2) / 3
var isEqual = (value == 2)
var left = false
var right = true
var LeftAndRight = left && right
var LeftOrRight = left || right
var LeftReverse = !left
var point = 5
var points = 9840.0
var pointss = 10
var fruits = [4]string{"apel", "jeruk", "durian", "pisang"}
var chicken map[string]int
var randomizer = rand.New(rand.NewSource(time.Now().Unix()))
var chickens = map[string]int{
	"januari":  50,
	"februari": 40,
	"maret":    34,
	"april":    67,
}

func main() {
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilanagan desimal: %.3f\n", decimalNumber)
	fmt.Printf("exist? %t \n", exist)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)
	fmt.Printf("left && right \t(%t) \n", LeftAndRight)
	fmt.Printf("left || right \t(%t) \n", LeftOrRight)
	fmt.Printf("!left \t\t(%t) \n", LeftReverse)
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}
	if percent := points / 100; percent >= 100 {
		fmt.Printf("%.1f%s sangat baik!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s baik\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s sangat buruk\n", percent, "%")
	}
	switch pointss {
	case 8, 9, 10:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
	for i := 0; i < 6; i++ {
		fmt.Println("angka", i)

	}
	var names [4]string
	names[0] = "syaiful"
	names[1] = "Imanudin"
	names[2] = "water"
	names[3] = "law"
	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println("jumlah buah \t\t", len(fruits))
	fmt.Println("nama buah \t", fruits)

	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) //januari 50
	fmt.Println("mei", chicken["mei"])         //mei 40

	for key, val := range chickens {
		fmt.Println(key, " \t:", val)
	}

	{
		var names = []string{"syaiful", "imanudin"}
		printMessage("woi", names)
	}
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -14)

	{
		var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
		var msg = fmt.Sprintf("rata-rata : %.2f", avg)
		fmt.Println(msg)
	}

	{
		var secret interface{}

		secret = "syaiful imanudin hunt"
		fmt.Println(secret)

		secret = []string{"apple", "manggo", "banana"}
		fmt.Println(secret)

		secret = 12.4
		fmt.Println(secret)
	}

	{
		var number = 21
		var reflectValue = reflect.ValueOf(number)

		fmt.Println("tipe variabel :", reflectValue.Type())

		if reflectValue.Kind() == reflect.Int {
			fmt.Println("nilai variabel :", reflectValue.Int())
		}

	}

	{
		var s1 = &student{Name: "Syaiful", Grade: 2}
		s1.getPropertyInfo()
	}

	{
		var input string
		fmt.Print("type some number: ")
		fmt.Scanln(&input)

		var number int
		var err error
		number, err = strconv.Atoi(input)

		if err == nil {
			fmt.Println(number, "is number")
		} else {
			fmt.Println(input, "is not number")
			fmt.Println(err.Error())
		}
	}

}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}

	{
		var s1 = &student{Name: "syaiful imanudin", Grade: 2}
		fmt.Println("nama :", s1.Name)

		var reflectValue = reflect.ValueOf(s1)
		var method = reflectValue.MethodByName("SetName")
		method.Call([]reflect.Value{
			reflect.ValueOf("kiki"),
		})
		fmt.Println("nama :", s1.Name)
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}
	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

const (
	square         = "kotak"
	isToday  bool  = true
	numeric  uint8 = 1
	floatNum       = 2.2
)

type student struct {
	Name  string
	Grade int
}
