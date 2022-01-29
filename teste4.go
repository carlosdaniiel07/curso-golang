package main

import (
	"fmt"
)

type Pedido struct {
	cliente string
	itens []ItemPedido
}
type ItemPedido struct {
	descricao string
	quantidade int
	valor float64
}

func (pedido Pedido) getValorTotal() float64 {
	total := 0.0

	for _, item := range pedido.itens {
		total += item.subTotal()
	}

	return total
}

func (item ItemPedido) subTotal() float64 {
	return item.valor * float64(item.quantidade)
}

func main() {
	pedido := Pedido {
		cliente: "Carlos",
		itens: []ItemPedido {
			ItemPedido {
				descricao: "Teste",
				quantidade: 1,
				valor: 200,
			},
			{
				descricao: "Teste 2",
				quantidade: 2,
				valor: 500,
			},
		},
	}

	fmt.Println(pedido)
	fmt.Printf("Valor total do pedido R$ %.2f", pedido.getValorTotal())
}
