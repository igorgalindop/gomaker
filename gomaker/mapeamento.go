package gomaker

// Retorna o objeto mapa para armazenar valores. Chave (string) e Valor (qualquer tipo)
func CriarObjetoMapa() map[interface{}]interface{} {
	return make(map[interface{}]interface{})
}

// MapeamentoExisteChave retorna se uma determinada chave existe no objeto mapa.
func MapeamentoExisteChave(m map[interface{}]interface{}, chave interface{}) bool {
	_, existe := m[chave]
	return existe
}

// Cria um mapa para armazenar valores a partir de uma lista de "par/valor".
func CriarObjetoMapaComParametros(pares ...[]interface{}) map[interface{}]interface{} {
	mapa := make(map[interface{}]interface{})

	for _, lista := range pares {
		chave := lista[0]
		valor := lista[1]
		mapa[chave] = valor
	}

	for i := 0; i < len(pares); i += 2 {

	}
	return mapa
}
