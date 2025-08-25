# Aprendendo Go

Ao executar o comando abaixo

```sh
go run . 123 456 789
```

a saida esperada eh

```text
args [0] /tmp/pastatemporaria/ch1
args ... 123 456 789
```

## com build tags

Ao executar o comando abaixo

```sh
go run -tags=vscode .
```

a saida esperada eh

```text
[go:build:vscode] rodando com -tags=vscode
[go:build:vscode] executando a funcao init
[go:build:vscode] do arquivo debug.go
args [0] /tmp/pastatemporaria/ch1
args ... 123 456 789
```
