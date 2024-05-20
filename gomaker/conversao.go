package gomaker

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
Tenta converter o valor para um número inteiro.
Se a conversão for bem-sucedida, retorna o número fracionado.
Se a conversão falhar, retorna um erro.
*/
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

/*
Tenta converter o valor para um número fracionado.
Se a conversão for bem-sucedida, retorna o número fracionado.
Se a conversão falhar, retorna um erro.
*/
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

// Letras converte um valor para uma string.
func ParaLetras(valor interface{}) string {
	return fmt.Sprint(valor)
}

/*
Converte para data o valor passado no primeiro parâmetro.
O segundo parâmetro é opcional, e por ele pode-se informar o formato da data que será retornada.

Parâmetros:

1. Valor a ser convertido

2. Formato do parâmetro a ser convertido. (opcional, fomato padrão dd/MM/yyyy)

Observações:

1. O formato da Data referente ao segundo parâmetro deve ser escrito no seguinte formato:

	    dd = Dia
		MM = Mês
		yyyy = Ano
		HH = Horas
		mm = Minutos
		ss = Segundos
*/
func ParaData(valor string, formato ...string) (time.Time, error) {
	defaultFormat := "02/01/2006" // dd/MM/yyyy

	var layout string

	if len(formato) > 0 && formato[0] != "" {
		layout = formato[0]
		layout = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(layout, "yyyy", "2006"), "dd", "02"), "MM", "01")
		layout = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(layout, "HH", "15"), "mm", "04"), "ss", "05")
	} else {

		layout = defaultFormat
	}

	valor = valor[:len(layout)]
	fmt.Println(valor)

	data, err := time.Parse(layout, valor)
	if err != nil {
		return time.Time{}, fmt.Errorf("erro ao converter para data: %v", err)
	}

	return data, nil
}
