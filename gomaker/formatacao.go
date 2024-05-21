package gomaker

import (
	"strings"
	"time"
)

// Função para formatar a data conforme o formato especificado
func FormatacaoDataHora(data time.Time, formato string) string {
	replacements := map[string]string{
		"yyyy": "2006",
		"MM":   "01",
		"dd":   "02",
		"H":    "15",
		"k":    "15",
		"K":    "3",
		"h":    "03",
		"a":    "PM",
		"mm":   "04",
		"ss":   "05",
		"SSS":  ".000",
	}

	for key, value := range replacements {
		formato = strings.Replace(formato, key, value, -1)
	}

	return data.Format(formato)
}
