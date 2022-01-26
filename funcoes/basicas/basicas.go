package main

import "fmt"

func F1() {
	fmt.Println("F1")
}

func F2(p1 string, p2 string) {
	fmt.Printf("F2: %s, %s\n", p1, p2)
}

func F2_v2(p1, p2 string)  {
	fmt.Printf("F2_v2: %s, %s\n", p1, p2)
}

func F3() string {
	return "F3"
}

func F4() (int, string) {
	return 2022, "Janeiro"
}

func main() {
	F1()
	F2("Carlos", "Daniel")
	F2_v2("Carlos", "Daniel")

	fmt.Println(F3())

	ano, mes := F4()

	fmt.Println(ano, mes)
}
