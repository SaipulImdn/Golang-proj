package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("halo")
	fmt.Println("selamat datang")

	{
		orderSomeFood("pizza")
		orderSomeFood("burger")
	}

	{
		number := 3

		if number == 3 {
			fmt.Println("halo 1")
			func() {
				defer fmt.Println("halo 3")
			}()
		}
		fmt.Println("halo 2")

	}

	{
		defer fmt.Println("halo")
		os.Exit(1)
		fmt.Println("selamat datang")
	}
}

func orderSomeFood(menu string) {
	defer fmt.Println("terimakasih, silahkan tunggu")
	if menu == "pizza" {
		fmt.Print("pilihan tepat!", "")
		fmt.Print("pizza diteampat kami paling enak!", "\n")
		return
	}

	fmt.Println("pesanan anda:", menu)
}
