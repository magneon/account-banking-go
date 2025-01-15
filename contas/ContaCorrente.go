package contas

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (conta *ContaCorrente) Withdraw(valor float64) {
	if valor > 0 {
		if conta.Saldo > valor {
			conta.Saldo -= valor
			fmt.Println("Saque de", valor, "realizado com sucesso da conta de", conta.Titular, ". Novo Saldo", conta.Saldo)
		} else {
			fmt.Println("Saldo insuficiente. Saldo atual", conta.Saldo)
		}
	}
}

func (conta *ContaCorrente) Deposit(valor float64) {
	if valor > 0 {
		conta.Saldo += valor
		fmt.Println("Depósito de", valor, "realizado com sucesso na conta de", conta.Titular, ". Novo Saldo", conta.Saldo)
	} else {
		fmt.Println("Não é possível depositar valores negativos. Refaça a operação.")
	}
}

func (origem *ContaCorrente) Transfer(valor float64, destino *ContaCorrente) {
	if valor > 0 {
		if origem.Saldo >= valor {
			destino.Deposit(valor)
			origem.Withdraw(valor)

			fmt.Println("Transferindo", valor, "de", origem.Titular, "para", destino.Titular)
		} else {
			fmt.Println("Saldo insuficiente para realizar essa transferência.")
		}
	} else {
		fmt.Println("Não é possível transferir valores negativos. Refaça a operação.")
	}
}
