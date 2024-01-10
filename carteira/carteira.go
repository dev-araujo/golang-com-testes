package main

import "fmt"

type Carteira struct {
	saldo int
}

func (c Carteira) Depositar(quantidade int) {
	fmt.Printf("O endereço do saldo Depositar é %v \v", &c.saldo)
	c.saldo += quantidade
}

func (c Carteira) Saldo() int {

	return saldo
}
