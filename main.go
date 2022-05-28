package main

import (
	"fmt"

	"github.com/nenodias/banco-golang/clientes"
	c "github.com/nenodias/banco-golang/contas"
)

func PagarBoleto(conta Conta, valor float64) {
	conta.Sacar(valor)
}

type Conta interface {
	Sacar(valor float64) (err error)
}

func main() {
	cliente := clientes.Titular{Nome: "Guilherme"}
	titular := c.ContaCorrente{Titular: cliente, NumeroAgencia: 589, NumeroConta: 123456}
	titular.Depositar(125.50)
	fmt.Println(titular)
	cliente2 := clientes.Titular{Nome: "Bruna"}
	titular2 := c.ContaPoupanca{Titular: cliente2, NumeroAgencia: 222, NumeroConta: 111222}
	titular2.Depositar(325.99)
	fmt.Println(titular2)
	err := titular2.Sacar(50)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(titular2)

	err = titular2.Depositar(150)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(titular2)

	err = titular2.Transferir(500, &titular)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(titular, titular2)

	PagarBoleto(&titular, 25.50)
	fmt.Println(titular.GetSaldo())

	PagarBoleto(&titular2, 25.99)
	fmt.Println(titular2.GetSaldo())

	/*	Exemplo de código utilizando ponteiros

		var contaCris *c.ContaCorrente = new(c.ContaCorrente)
		contaCris.Titular = "Cris"
		fmt.Println(*contaCris)
	*/
}
