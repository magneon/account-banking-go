package contas

import (
	"account-banking-go/titulares"
	"fmt"
)

type ContaCorrente struct {
	Titular       titulares.PessoaFisica
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Withdraw(valor float64) {
	if valor > 0 {
		if conta.saldo > valor {
			conta.saldo -= valor
			fmt.Println("Saque de", valor, "realizado com sucesso da conta de", conta.Titular.Nome, ". Novo Saldo", conta.saldo)
		} else {
			fmt.Println("Saldo insuficiente. Saldo atual", conta.saldo)
		}
	}
}

func (conta *ContaCorrente) Deposit(valor float64) {
	if valor > 0 {
		conta.saldo += valor
		fmt.Println("Depósito de", valor, "realizado com sucesso na conta de", conta.Titular.Nome, ". Novo Saldo", conta.saldo)
	} else {
		fmt.Println("Não é possível depositar valores negativos. Refaça a operação.")
	}
}

func (origem *ContaCorrente) Transfer(valor float64, destino *ContaCorrente) {
	if valor > 0 {
		if origem.saldo >= valor {
			destino.Deposit(valor)
			origem.Withdraw(valor)

			fmt.Println("Transferindo", valor, "de", origem.Titular.Nome, "para", destino.Titular.Nome)
		} else {
			fmt.Println("Saldo insuficiente para realizar essa transferência.")
		}
	} else {
		fmt.Println("Não é possível transferir valores negativos. Refaça a operação.")
	}
}

func (conta *ContaCorrente) Extract() {
	fmt.Println("Seu saldo atualmente é de R$", conta.saldo)
}
