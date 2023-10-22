package enderecos

import "testing"

type enderecoDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []enderecoDeTeste{
		{"Rua Cebola", "Rua"},
		{"Avenida Cebola", "Avenida"},
		{"Rodovia Cebola", "Rodovia"},
		{"Estrada Cebola", "Estrada"},
		// {"", "Tipo invalido"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoEsperado := TipoDeEndereco(cenario.enderecoInserido)
		if retornoEsperado != cenario.retornoEsperado {
			t.Errorf("Tipo recebido inesperado, esperava %s e recebeu %s", retornoEsperado, cenario.retornoEsperado)
		}
	}
}
