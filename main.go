package main

import (
	contas "account-banking-go/contas"
	"fmt"
)

func main() {
	fmt.Println("Olá mundo, Go!")

	var titular string = "Rafael de Mesquita Moura"
	var numeroAgencia int = 1
	var numeroConta int = 40576891
	var saldo float64 = 0.0

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	conta1 := contas.ContaCorrente{Titular: titular, NumeroAgencia: numeroAgencia, NumeroConta: numeroConta, Saldo: saldo}
	fmt.Println(conta1)

	conta2 := contas.ContaCorrente{Titular: "Talita de Lima Mesquita", NumeroAgencia: 1, NumeroConta: 12345, Saldo: 1000.0}
	fmt.Println(conta2)

	var conta3 *contas.ContaCorrente
	conta3 = new(contas.ContaCorrente)
	conta3.Titular = "Francisca Francineide de Mesquita Moura"
	conta3.Saldo = 16000.0
	fmt.Println(conta3)

	// conta3.Withdraw(1100)
	// conta3.Deposit(-10000)
	conta3.Transfer(1500, &conta1)
}
