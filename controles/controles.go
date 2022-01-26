package main

import (
	"fmt"
	"math/rand"
	"time"
)

// não há parenteses; sempre precisa ter chaves para delimitar bloco;
func imprimirResultado(nota int) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else if nota == 5 {
		fmt.Println("Recuperação!")
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

// estrutura switch; não é necessário usar o "break" assim como em outras linguagens
func notaParaConceito(nota int) string {
	switch nota {
		case 10:
			fmt.Println("Case 10")

			// desce para o próximo case (case 9)
			fallthrough
		case 9:
			return "A"
		case 8, 7:
			return "B"
		case 6, 5:
			return "C"
		case 4, 3:
			return "D"
		case 2, 1, 0:
			return "E"
		default:
			return "Nota inválida!"
	}
}

func numeroAleatorio() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	return random.Intn(10)
}

// é a única estrutura de repetição do Go
func numerosPares(valorMaximo int) {
	contador := 0

	for contador <= valorMaximo {
		fmt.Print(contador, ",")
		contador += 2
	}
}

// é a única estrutura de repetição do Go
func numerosImpares(valorMaximo int)  {
	for i := 1; i <= valorMaximo; i += 2 {
		fmt.Print(i, ",")
	}
}

// laço infinito
func lacoInfinito()  {
	for true {
		fmt.Println("Laço infinito...")
		time.Sleep(time.Second * 3)
	}
}

func main()  {
	imprimirResultado(7)
	imprimirResultado(5)
	imprimirResultado(3)

	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(7))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(3))

	// iniciando variável no bloco if, "if init"
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu")
	}

	numerosPares(12)
	fmt.Println()
	numerosImpares(12)
	fmt.Println()

	lacoInfinito()
}
