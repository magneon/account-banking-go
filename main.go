package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	fmt.Println("Olá mundo, Go!")

	var titular string = "Rafael de Mesquita Moura"
	var numeroAgencia int = 1
	var numeroConta int = 40576891
	var saldo float64 = 0.0

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	conta1 := ContaCorrente{titular, numeroAgencia, numeroConta, saldo}
	fmt.Println(conta1)

	conta2 := ContaCorrente{titular: "Talita de Lima Mesquita", numeroAgencia: 1, numeroConta: 12345, saldo: 1000.0}
	fmt.Println(conta2)

	var conta3 *ContaCorrente
	conta3 = new(ContaCorrente)
	conta3.titular = "Francisca Francineide de Mesquita Moura"
	conta3.saldo = 16000.0
	fmt.Println(conta3)

	conta3.Withdraw(1100)
	conta3.Deposit(-10000)
}

func (conta *ContaCorrente) Withdraw(valor float64) {
	if valor > 0 {
		if conta.saldo > valor {
			conta.saldo -= valor
			fmt.Println("Saque de", valor, "realizado com sucesso da conta de", conta.titular, ". Novo saldo", conta.saldo)
		} else {
			fmt.Println("Saldo insuficiente. Saldo atual", conta.saldo)
		}
	}
}

func (conta *ContaCorrente) Deposit(valor float64) {
	if valor > 0 {
		conta.saldo += valor
		fmt.Println("Depósito de", valor, "realizado com sucesso na conta de", conta.titular, ". Novo saldo", conta.saldo)
	} else {
		fmt.Println("Não é possível depositar valores negativos. Refaça a operação.")
	}
}
