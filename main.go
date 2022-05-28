package main

import (
	"fmt"

	c "github.com/nenodias/banco-golang/contas"
)

func main() {
	titular := c.ContaCorrente{Titular: "Guilherme", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}
	fmt.Println(titular)
	titular2 := c.ContaCorrente{"Bruna", 222, 111222, 200}
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

	/*	Exemplo de código utilizando ponteiros

		var contaCris *c.ContaCorrente = new(c.ContaCorrente)
		contaCris.Titular = "Cris"
		fmt.Println(*contaCris)
	*/
}
