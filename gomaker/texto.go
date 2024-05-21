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

/*
Esta função localiza todas as subseqüências iguais ao 2º parâmetro dentro do texto e os substituem pelo conteúdo passado no
3º parâmetro.
*/
func TrocarTodasAsSubsequenciasRegex(texto, valorAEncontrar, valorDesejado string) (string, error) {
	regex, err := regexp.Compile(valorAEncontrar)
	if err != nil {
		return "", fmt.Errorf("erro ao compilar regex: %v", err)
	}

	resultado := regex.ReplaceAllString(texto, valorDesejado)
	return resultado, nil
}

// Procura dentro do texto a subsequência passada como parâmetro e retorna verdadeiro se encontrada.
func ExisteSubsequencia(texto, valorALocalizar string) bool {
	return strings.Contains(texto, valorALocalizar)
}

/*
A função localiza o 2º parâmetro dentro do texto e o substitui pelo conteúdo passado no 3º parâmetro.

Observação:

A função só substitui a primeira ocorrência encontrada. Para substituir todas as subsequências encontradas utilize a função
TrocarTodasAsSubsequencias.
*/
func TrocarSubsequencia(texto, valorAEncontrar, valorDesejado string) string {
	return strings.Replace(texto, valorAEncontrar, valorDesejado, 1)
}

/*
Esta função localiza todas as subseqüências iguais ao 2º parâmetro dentro do texto e os substituem pelo conteúdo passado no
3º parâmetro.
*/
func TrocarTodasAsSubsequencias(texto, valorAEncontrar, valorDesejado string) string {
	return strings.ReplaceAll(texto, valorAEncontrar, valorDesejado)
}

/*
IndiceSubsequencia localiza o conteúdo do 2° parâmetro dentro do conteúdo do primeiro parâmetro
e retorna a posição deste caso o encontre. Caso não encontre, retorna 0.
*/
func IndiceSubsequencia(texto, valorAEncontrar string) int {
	if valorAEncontrar == "" {
		return 0
	}
	indice := strings.Index(texto, valorAEncontrar)
	if indice == -1 {
		return 0
	}
	return indice + 1
}

// Faz uma validação numa string utilizando expressão regular.
func ValidarTextoUtilizandoExpressaoRegular(texto, expressaoRegular string) bool {
	padrao, err := regexp.Compile(expressaoRegular)
	if err != nil {
		return false
	}
	return padrao.MatchString(texto)
}

/*
Recebe um texto e retorna a quantidade de caracteres passada
como segundo parâmetro a partir do final.
*/
func SubsequenciaInvertida(texto string, quantidade int) string {
	runes := []rune(texto)
	if quantidade > len(runes) {
		quantidade = len(runes)
	}
	inicio := len(runes) - quantidade
	return string(runes[inicio:])
}

// Remove os caracteres das posições indicadas entre o segundo e terceiro parâmetro do texto passado no primeiro parâmetro.
func RemoverSubsequencia(texto string, posicaoInicial, quantidade int) string {
	runes := []rune(texto)
	inicio := posicaoInicial - 1
	fim := inicio + quantidade

	if inicio < 0 {
		inicio = 0
	}
	if fim > len(runes) {
		fim = len(runes)
	}
	if inicio >= len(runes) || inicio >= fim {
		return texto
	}

	return string(runes[:inicio]) + string(runes[fim:])
}

// Retorna a subsequência contida no texto que inicia na posição especificada e tem o tamanho indicado.
func Subsequencia(texto string, posicaoInicial, tamanho int) string {
	runes := []rune(texto)
	inicio := posicaoInicial - 1
	fim := inicio + tamanho

	if inicio < 0 {
		inicio = 0
	}
	if fim > len(runes) {
		fim = len(runes)
	}
	if inicio >= len(runes) || inicio >= fim {
		return ""
	}

	return string(runes[inicio:fim])
}

// Completa o texto à esquerda com o conteúdo fornecido até atingir o tamanho especificado.
func CompletarAEsquerda(texto string, tamanho int, conteudo string) string {
	repeticoes := tamanho - len(texto)
	if repeticoes <= 0 {
		return texto
	}

	return strings.Repeat(conteudo, repeticoes) + texto
}

// Completa o texto à direita com o conteúdo fornecido até atingir o tamanho especificado.
func CompletarADireita(texto string, tamanho int, conteudo string) string {
	repeticoes := tamanho - len(texto)
	if repeticoes <= 0 {
		return texto
	}

	return texto + strings.Repeat(conteudo, repeticoes)
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
