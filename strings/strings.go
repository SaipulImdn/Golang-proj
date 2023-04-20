package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExits = strings.Contains("Syaiful imanudin", "kiki")
	fmt.Println(isExits)

	{
		var howMany = strings.Count("Syaiful Imanudin", "i")
		fmt.Println(howMany)
	}

	{
		var text = "banana"
		var find = "a"
		var replaceWith = "o"

		var newText1 = strings.Replace(text, find, replaceWith, 1)
		fmt.Println(newText1)
		var newText2 = strings.Replace(text, find, replaceWith, 2)
		fmt.Println(newText2)
		var newText3 = strings.Replace(text, find, replaceWith, -1)
		fmt.Println(newText3)
	}
}
