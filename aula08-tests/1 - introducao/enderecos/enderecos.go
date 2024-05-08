package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEnderco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEnderco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(primeiraPalavraEnderco)
	}

	return "Tipo Inválido"
}
