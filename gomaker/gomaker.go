package gomaker

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

// Retorna uma lista vazia que pode ser populada com valores de qualquer tipo.
func CriarObjetoLista() []interface{} {
	return []interface{}{}
}

// Retorna o objeto mapa para armazenar valores. Chave (string) e Valor (qualquer tipo)
func CriarObjetoMapa() map[string]interface{} {
	return make(map[string]interface{})
}

// Insere um elemento em uma lista na posição indicada (ou no final).
func InserirObjetoNaLista(lista *[]interface{}, elemento interface{}, posicao *int) {
	if posicao == nil || *posicao >= len(*lista) {
		*lista = append(*lista, elemento)
		return
	}

	*lista = append((*lista)[:*posicao], append([]interface{}{elemento}, (*lista)[*posicao:]...)...)
}

// Retorna o MD5 de um texto como uma string hexadecimal.
func MD5doTexto(texto string) string {
	hash := md5.Sum([]byte(texto))
	return hex.EncodeToString(hash[:])
}

// Junta vários itens de texto em apenas um item. Retorna o valor concatenado.
func Append(texto *string, variants ...interface{}) string {
	var strVariants []string
	for _, v := range variants {
		strVariants = append(strVariants, fmt.Sprint(v))
	}
	*texto += strings.Join(strVariants, "")
	return *texto
}

// Remove todos os espaços de uma string.
func RemoverEspacos(texto string) string {
	return strings.ReplaceAll(texto, " ", "")
}

// ENuloOuVazio verifica se um valor é nulo ou vazio.
func ENuloOuVazio(valor interface{}) bool {
	if valor == nil {
		return true
	}

	switch v := reflect.ValueOf(valor); v.Kind() {
	case reflect.String:
		return strings.TrimSpace(v.String()) == ""
	case reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Struct:
		return reflect.DeepEqual(valor, reflect.Zero(reflect.TypeOf(valor)).Interface())
	default:
		return false
	}
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
