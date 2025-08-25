//go:build vscode

package main

import "fmt"

func init() {
	fmt.Println("[go:tags:vscode] rodando com -tags=vscode")
	fmt.Println("[go:tags:vscode] executando a funcao init")
	fmt.Println("[go:tags:vscode] do arquivo debug.go")
}
