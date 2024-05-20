package gomaker

// Retorna uma lista vazia que pode ser populada com valores de qualquer tipo.
func CriarObjetoLista() []interface{} {
	return []interface{}{}
}

// Cria uma lista com todos os valores que foram passados pelos parâmetros.
func CriarListaAPartirDosElementos(elementos ...interface{}) []interface{} {
	return elementos
}

// Insere um elemento em uma lista na posição indicada (ou no final).
func InserirObjetoNaLista(lista *[]interface{}, elemento interface{}, posicao *int) {
	if posicao == nil || *posicao >= len(*lista) {
		*lista = append(*lista, elemento)
		return
	}

	*lista = append((*lista)[:*posicao], append([]interface{}{elemento}, (*lista)[*posicao:]...)...)
}
