# Go CNPJ [![Tests](https://github.com/cuducos/go-cnpj/workflows/Tests/badge.svg)]()[![GoDoc](https://godoc.org/github.com/cuducos/go-cnpj?status.svg)](https://godoc.org/github.com/cuducos/go-cnpj)![Go version](https://img.shields.io/github/go-mod/go-version/cuducos/go-cnpj)

A Go module to validate CNPJ numbers (Brazilian companies' unique identifier for the Federal Revenue).

> [!IMPORTANT]
> Starting in July 2026 [the CNPJ number will be alphanumeric](https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/cnpj-alfanumerico). This package **already supports the new format**. If you **do not** want to support the new format, tag this package to [`v0.1.1`](https://github.com/cuducos/go-cnpj/releases/tag/v0.1.1).

```go
package main

import "github.com/cuducos/go-cnpj"


func main() {
	// these return true
	cnpj.IsValid("11222333000181")
	cnpj.IsValid("11.222.333/0001-81")
	cnpj.IsValid("12.ABC.345/01DE-35")
	cnpj.IsValid("12ABC34501DE35")

	// these return false
	cnpj.IsValid("11.111.111/1111-11")
	cnpj.IsValid("12.345.678 9012-34")
	cnpj.IsValid("AB.CDE.FGH/IJKL-MN")
	cnpj.IsValid("123")

	// these return 11111111111111 and 12ABC34501DE35
	cnpj.Unmask("11.111.111/1111-11")
	cnpj.Unmask("12.ABC.345/01DE-35")

	// these return 11.111.111/1111-11 and 12.ABC.345/01DE-35
	cnpj.Mask("11111111111111")
	cnpj.Mask("12ABC34501DE35")
}
```

Based on [Go CPF](https://github.com/cuducos/go-cpf) ❤️
