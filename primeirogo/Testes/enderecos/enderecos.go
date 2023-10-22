package enderecos

import "strings"

// Funcao TipoDeEndereco :)
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "estrada", "avenida", "rodovia"}

	value := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(value, " ")[0]

	enderecoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo invalido"

}
