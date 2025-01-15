package main

import (
	contas "account-banking-go/contas"
	titulares "account-banking-go/titulares"
	"fmt"
)

func main() {
	fmt.Println("Olá mundo, Go!")

	var titular string = "Rafael de Mesquita Moura"
	var cpf = "37361310893"
	var profissao = "Desenvolvedor"
	var numeroAgencia int = 1
	var numeroConta int = 40576891
	var saldo float64 = 0.0
	// fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	clienteRafael := titulares.PessoaFisica{Nome: titular, CPF: cpf, Profissao: profissao}
	conta1 := contas.ContaCorrente{Titular: clienteRafael, NumeroAgencia: numeroAgencia, NumeroConta: numeroConta}
	conta1.Deposit(500)
	// fmt.Println(conta1)

	clienteTalita := titulares.PessoaFisica{Nome: "Talita de Lima Mesquita", CPF: "39133388881", Profissao: "Confeiteira"}
	conta2 := contas.ContaCorrente{Titular: clienteTalita, NumeroAgencia: 1, NumeroConta: 12345}
	conta2.Deposit(500)
	// fmt.Println(conta2)

	var conta3 *contas.ContaCorrente = new(contas.ContaCorrente)
	conta3.Titular.Nome = "Francisca Francineide de Mesquita Moura"
	conta3.Deposit(16000.0)
	// fmt.Println(conta3)

	// conta3.Withdraw(1100)
	// conta3.Deposit(-10000)
	conta3.Transfer(1500, &conta1)

	conta1.Extract()
	conta2.Extract()
	conta3.Extract()
}
