package gomaker

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// Verifica se o caminho especificado existe.
func ExistePasta(caminho string) bool {
	caminhoAbsoluto, err := filepath.Abs(caminho)
	if err != nil {
		fmt.Println("Erro ao obter caminho absoluto:", err)
		return false
	}

	if _, err := os.Stat(caminhoAbsoluto); os.IsNotExist(err) {
		return false
	}
	return true
}

/*
Cria uma nova pasta no servidor.

Observações:

 1. Caso a pasta definida no parâmetro já exista, a funçao não retornará erro.
*/
func CriarNovaPasta(caminho string) error {
	if _, err := os.Stat(caminho); err == nil {
		return nil
	}

	if err := os.MkdirAll(caminho, 0755); err != nil {
		return fmt.Errorf("erro ao criar a pasta '%s': %v", caminho, err)
	}

	return nil
}

/*
Abre um arquivo para edição e retorna uma referência para o arquivo.

Parâmetros:

 1. Caminho do arquivo.
 2. Falso, para abrir o arquivo para edição sobrepondo o seu conteúdo original, ou verdadeiro, para abrir e apenas inserir novos dados ao conteúdo existente.

Observação:

 1. Caso o arquivo não exista, ele será criado automaticamente.
*/
func AbrirArquivoParaEscrita(caminho string, adicionar bool) (*os.File, error) {
	var modoAbertura int
	if adicionar {
		modoAbertura = os.O_WRONLY | os.O_APPEND | os.O_CREATE
	} else {
		modoAbertura = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	}

	arquivo, err := os.OpenFile(caminho, modoAbertura, 0644)
	if err != nil {
		return nil, err
	}

	return arquivo, nil
}

/*
Força a gravação do arquivo em disco, movendo todo o conteúdo que estava na memória para a unidade de
armazenamento. O uso desta função evita perdas inesperadas e força o Sistema Operacional  a gravar logo
aquelas informações(contidas na memória) no arquivo.
*/
func DescarregarParaArquivo(arquivo *os.File) error {
	if arquivo == nil {
		return fmt.Errorf("arquivo não pode ser nulo")
	}

	if _, err := arquivo.Seek(0, 0); err != nil {
		return err
	}

	if err := arquivo.Sync(); err != nil {
		return err
	}

	return nil
}

// Fecha um arquivo que foi aberto para escrita ou leitura.
func FecharArquivo(arquivo *os.File) error {
	if arquivo == nil {
		return fmt.Errorf("arquivo não pode ser nulo")
	}

	if err := arquivo.Sync(); err != nil {
		return err
	}

	if err := arquivo.Close(); err != nil {
		return err
	}

	return nil
}

/*
AdicionarAoArquivo adiciona conteúdo ao arquivo especificado.
Armazena no arquivo o conteúdo passado pelo segundo parâmetro.

Parâmetros:

 1. Variável onde a referência do arquivo foi armazenada. Essa referência é o retorno da função "Abrir arquivo para escrita".
 2. Conteúdo que será adicionado ao arquivo.
 3. Charset (Ex.: UTF-8, ISO-8859-1)

Observações

 1. No terceiro parâmetro, deverá ser informado o charset do arquivo. Caso não seja passado o charset, a função adotará o charset UTF8.
*/
func AdicionarAoArquivo(arquivo *os.File, conteudo, charset string) error {
	if arquivo == nil {
		return fmt.Errorf("arquivo não pode ser nulo")
	}

	var writer io.Writer

	switch strings.ToUpper(charset) {
	case "UTF-8":
		writer = arquivo
	case "ISO-8859-1":
		writer = transform.NewWriter(arquivo, charmap.ISO8859_1.NewEncoder())
	default:
		writer = transform.NewWriter(arquivo, unicode.UTF8.NewEncoder())
	}

	if _, err := writer.Write([]byte(conteudo)); err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}

	return nil
}
