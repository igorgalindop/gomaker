package gomaker

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// Ano retorna o ano como um inteiro de acordo com a data informada.
func Ano(data time.Time) int {
	return data.Year()
}

// Dia retorna o dia como um inteiro de acordo com a data informada.
func Dia(data time.Time) int {
	return data.Day()
}

// Mes retorna o mês como um inteiro de acordo com a data informada.
func Mes(data time.Time) int {
	return int(data.Month())
}

/*
Timezones

	0 = Padrão (de acordo com o timezone do servidor)
	1 = America/Sao_Paulo
*/
func Hoje(timezone int) (time.Time, error) {

	if timezone != 0 {
		var timeLocation string
		var err error

		switch timezone {
		case 1:
			timeLocation = "America/Sao_Paulo"
		default:
			timeLocation = "America/Sao_Paulo"
		}

		timezone, err := time.LoadLocation(timeLocation)
		if err != nil {
			fmt.Println("Erro ao carregar o fuso horário:", err)
			return time.Now().In(timezone), err
		}

		return time.Now().In(timezone), nil
	}
	return time.Now(), nil
}

func CriarObjetoLista() []interface{} {
	return []interface{}{}
}

func CriarObjetoMapa() map[string]interface{} {
	return make(map[string]interface{})
}

func ParaInteiro(valor interface{}) (int64, error) {
	switch v := valor.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case string:
		// Tentar converter a string para float64
		floatValor, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return int64(floatValor), nil
	default:
		return 0, fmt.Errorf("valor não pode ser convertido para inteiro")
	}
}

// ParaFracionado tenta converter o valor para um número fracionado.
// Se a conversão for bem-sucedida, retorna o número fracionado.
// Se a conversão falhar, retorna um erro.
func ParaFracionado(valor interface{}) (float64, error) {
	switch v := valor.(type) {
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		// Tentar converter a string para float64
		resultado, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return resultado, nil
	default:
		return 0, fmt.Errorf("valor não pode ser convertido para fracionado")
	}
}

// InserirObjetoNaLista insere um elemento em uma lista na posição indicada (ou no final).
func InserirObjetoNaLista(lista *[]interface{}, elemento interface{}, posicao *int) {
	if posicao == nil || *posicao >= len(*lista) {
		// Se a posição for nula ou maior ou igual ao tamanho da lista,
		// insere o elemento no final da lista.
		*lista = append(*lista, elemento)
		return
	}

	// Caso contrário, insere o elemento na posição indicada.
	*lista = append((*lista)[:*posicao], append([]interface{}{elemento}, (*lista)[*posicao:]...)...)
}

// MD5doTexto retorna o MD5 de um texto como uma string hexadecimal.
func MD5doTexto(texto string) string {
	hash := md5.Sum([]byte(texto))
	return hex.EncodeToString(hash[:])
}
