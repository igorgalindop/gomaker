package gomaker

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Junta vários itens de texto em apenas um item. Retorna o valor concatenado.
func Append(texto *string, variants ...interface{}) string {
	var strVariants []string
	for _, v := range variants {
		strVariants = append(strVariants, fmt.Sprint(v))
	}
	*texto += strings.Join(strVariants, "")
	return *texto
}

// Retorna o MD5 de um texto como uma string hexadecimal.
func MD5doTexto(texto string) string {
	hash := md5.Sum([]byte(texto))
	return hex.EncodeToString(hash[:])
}

// Remove todos os espaços de uma string.
func RemoverEspacos(texto string) string {
	return strings.ReplaceAll(texto, " ", "")
}

// Converte uma string para maiúsculo.
func Maiusculo(texto string) string {
	return strings.ToUpper(texto)
}

// Converte uma string para minúsculo.
func Minusculo(texto string) string {
	return strings.ToLower(texto)
}

// Verifica se o conteúdo do primeiro parâmetro inicia com o conteúdo do segundo parâmetro.
func IniciaCom(texto, valorInicial string) bool {
	return strings.HasPrefix(texto, valorInicial)
}

// Retorna o símbolo correspondente na tabela ASCII do número passado como parâmetro.
func Letra(numero int) string {
	return fmt.Sprintf("%c", numero)
}

// Remove os acentos do texto passado por parâmetro
func RemoverAcentos(input string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, input)
	if err != nil {
		return input, err
	}
	return output, nil

}

// Retorna os caracteres deixando somente os números de um texto.
func SomenteNumeros(texto string) string {
	// Define a expressão regular para encontrar apenas números
	reg := regexp.MustCompile("[0-9]+")

	// Encontra todos os números no texto
	numeros := reg.FindAllString(texto, -1)

	// Junta todos os números encontrados em uma única string
	resultado := ""
	for _, numero := range numeros {
		resultado += numero
	}

	return resultado
}
