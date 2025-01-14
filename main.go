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
}
