package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)
	// -> []string{".../bab45", "banana", "potato", "ice cream"}

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)
	// -> []string{"banana", "potatao", "ice cream"}

	{
		var name = flag.String("name", "anonymous", "type your name")
		var age = flag.Int64("age", 21, "type your age")

		flag.Parse()
		fmt.Printf("name\t %s\n", *name)
		fmt.Printf("age\t: %d\n", *age)
	}
}
