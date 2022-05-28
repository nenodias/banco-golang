package contas

import (
	"errors"
	"fmt"

	"github.com/nenodias/banco-golang/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) (err error) {
	if valor <= c.saldo && valor > 0 {
		c.saldo -= valor
		return nil
	}
	msg := fmt.Sprintf("Erro ao sacar o valor de: %.2f (saldo=%.2f)", valor, c.saldo)
	return errors.New(msg)
}

func (c *ContaPoupanca) Depositar(valor float64) (err error) {
	if valor > 0 {
		c.saldo += valor
		return nil
	}
	msg := fmt.Sprintf("Valor %.2f inválido para depósito", valor)
	return errors.New(msg)
}

func (c *ContaPoupanca) Transferir(valor float64, contaDestino *ContaCorrente) (err error) {
	if valor <= c.saldo && valor > 0 {
		c.saldo -= valor
		err := contaDestino.Depositar(valor)
		if err != nil {
			c.saldo += valor
			return err
		}
		return nil
	}
	msg := fmt.Sprintf("Erro ao transferir o valor de: %.2f (saldo=%.2f)", valor, c.saldo)
	return errors.New(msg)
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
