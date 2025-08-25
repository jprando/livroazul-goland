//go:build vscode

package main

import "fmt"

func init() {
	fmt.Println("[go:build:vscode] rodando com -tags=vscode")
	fmt.Println("[go:build:vscode] executando a funcao init")
	fmt.Println("[go:build:vscode] do arquivo debug.go")
}
