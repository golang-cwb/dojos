package dojo

import "testing"

var casosDeTestes = []struct {
	titulo string
	input  int
	output string
}{
	{
		titulo: "Valor divisivel por 100",
		input:  300,
		output: "3 x R$ 100",
	},
	{
		titulo: "Valor divisivel por 100",
		input:  200,
		output: "2 x R$ 100",
	},
	{
		titulo: "Valor divisivel por 50",
		input:  50,
		output: "1 x R$ 50",
	},
	{
		titulo: "Valor divisivel por 20",
		input:  40,
		output: "2 x R$ 20",
	},
	{
		titulo: "Valor divisível por 10",
		input:  10,
		output: "1 x R$ 10",
	},
	{
		titulo: "Valor divisível por 100 e 20",
		input:  120,
		output: "1 x R$ 100 - 1 x R$ 20",
	},
	{
		titulo: "Valor sem notas disponiveis",
		input:  3,
		output: "O sistema nao possui notas disponiveis",
	},
}

func TestSaque(t *testing.T) {
	for _, teste := range casosDeTestes {
		t.Run(teste.titulo, func(t *testing.T) {
			out := saque(teste.input)
			if out != teste.output {
				t.Errorf("Erro ao calcular número de notas: Esperado %q, obtido %q", teste.output, out)
			}
		})
	}
}
