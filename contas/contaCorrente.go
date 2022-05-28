package contas

import (
	"errors"
	"fmt"
)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) (err error) {
	if valor <= c.Saldo && valor > 0 {
		c.Saldo -= valor
		return nil
	}
	msg := fmt.Sprintf("Erro ao sacar o valor de: %.2f (Saldo=%.2f)", valor, c.Saldo)
	return errors.New(msg)
}

func (c *ContaCorrente) Depositar(valor float64) (err error) {
	if valor > 0 {
		c.Saldo += valor
		return nil
	}
	msg := fmt.Sprintf("Valor %.2f inválido para depósito", valor)
	return errors.New(msg)
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) (err error) {
	if valor <= c.Saldo && valor > 0 {
		c.Saldo -= valor
		err := contaDestino.Depositar(valor)
		if err != nil {
			c.Saldo += valor
			return err
		}
		return nil
	}
	msg := fmt.Sprintf("Erro ao transferir o valor de: %.2f (Saldo=%.2f)", valor, c.Saldo)
	return errors.New(msg)
}
