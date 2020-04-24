package dojo

import (
	"fmt"
	"strings"
)

// Desenvolva um programa que simule a entrega de notas quando um cliente efetuar um saque em um caixa eletrônico. Os requisitos básicos são os seguintes:

// Entregar o menor número de notas;
// É possível sacar o valor solicitado com as notas disponíveis;
// Saldo do cliente infinito;
// Quantidade de notas infinito (pode-se colocar um valor finito de cédulas para aumentar a dificuldade do problema);
// Notas disponíveis de R$ 100,00; R$ 50,00; R$ 20,00 e R$ 10,00

func saque(valor int) string {
	cem := valor / 100
	var resultado []string
	if cem > 0 {
		resultado = append(resultado, fmt.Sprintf("%d x R$ 100", cem))
	}
	resto := valor % 100
	cinquenta := resto / 50
	if cinquenta > 0 {
		resultado = append(resultado, fmt.Sprintf("%d x R$ 50", cinquenta))
	}
	resto = resto % 50
	vinte := resto / 20
	if vinte > 0 {
		resultado = append(resultado, fmt.Sprintf("%d x R$ 20", vinte))
	}
	resto = resto % 20
	dez := resto / 10
	if dez > 0 {
		resultado = append(resultado, fmt.Sprintf("%d x R$ 10", dez))
	}
	if resto > 0 && resto < 10 {
		return "O sistema nao possui notas disponiveis"
	}
	return strings.Join(resultado, " - ")
}

func main() {
	fmt.Print("Hello Golang CWB")
}
