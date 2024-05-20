package gomaker

import (
	"reflect"
	"strings"
)

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
