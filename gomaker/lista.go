package gomaker

// Retorna uma lista vazia que pode ser populada com valores de qualquer tipo.
func CriarObjetoLista() []interface{} {
	return []interface{}{}
}

// Cria uma lista com todos os valores que foram passados pelos parâmetros.
func CriarListaAPartirDosElementos(elementos ...interface{}) []interface{} {
	return elementos
}
