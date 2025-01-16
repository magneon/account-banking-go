package contas

type Conta interface {
	PayBankSlip(valor float64)
}

func PayBankSlip(conta Conta, valor float64) {
	conta.PayBankSlip(valor)
}
