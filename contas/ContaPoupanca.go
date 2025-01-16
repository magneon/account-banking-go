package contas

import (
	"account-banking-go/titulares"
	"fmt"
)

type ContaPoupanca struct {
	Titular       titulares.PessoaFisica
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
	Operacao      int
}

func (conta *ContaPoupanca) Withdraw(valor float64) {
	if valor > 0 {
		if conta.saldo > valor {
			conta.saldo -= valor
			fmt.Println("Saque de", valor, "realizado com sucesso da conta de", conta.Titular.Nome, ". Novo Saldo", conta.saldo)
		} else {
			fmt.Println("Saldo insuficiente. Saldo atual", conta.saldo)
		}
	}
}

func (conta *ContaPoupanca) Deposit(valor float64) {
	if valor > 0 {
		conta.saldo += valor
		fmt.Println("Depósito de", valor, "realizado com sucesso na conta de", conta.Titular.Nome, ". Novo Saldo", conta.saldo)
	} else {
		fmt.Println("Não é possível depositar valores negativos. Refaça a operação.")
	}
}

func (conta *ContaPoupanca) Extract() {
	fmt.Println("CP: Seu saldo atualmente é de R$", conta.saldo)
}

func (conta *ContaPoupanca) PayBankSlip(valor float64) {
	if conta.saldo > 1000 && conta.saldo >= valor && valor > 0 {
		conta.saldo -= valor
		fmt.Println("Boleto pago com sucesso. Novo saldo", conta.saldo)
	} else {
		fmt.Println("Saldo insuficiente para pagar esse boleto")
	}
}
