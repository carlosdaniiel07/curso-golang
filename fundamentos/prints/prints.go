package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("Linha.")

	fmt.Println(" Nova linha")
	fmt.Println("olá")

	x := 3.141516
	xs := fmt.Sprint(x)

	fmt.Println("O valor de X é", x)
	fmt.Println("O valor de X é " + xs)
	fmt.Println("o valor de X é " + fmt.Sprint(x))

	fmt.Printf("O valor de X é %f", x)
	fmt.Println("")
	fmt.Printf("O valor de X é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "olá"

	fmt.Printf("\n%d - %.1f - %t - %s", a, b, c, d)
}
