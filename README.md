# Go CNPJ [![Tests](https://github.com/cuducos/go-cnpj/workflows/Tests/badge.svg)]()[![GoDoc](https://godoc.org/github.com/cuducos/go-cnpj?status.svg)](https://godoc.org/github.com/cuducos/go-cnpj)![Go version](https://img.shields.io/github/go-mod/go-version/cuducos/go-cnpj)

A Go module to validate CNPJ numbers (Brazilian companies unique identifier for the Federal Revenue).

```go
package main

import "github.com/cuducos/go-cnpj"


func main() {
	// these return true
	cnpj.IsValid("11222333000181")
	cnpj.IsValid("11.222.333/0001-81")

	// these return false
	cnpj.IsValid("11.111.111/1111-11")
	cnpj.IsValid("12.345.678 9012-34")
	cnpj.IsValid("AB.CDE.FGH/IJKL-MN")
	cnpj.IsValid("123")

	// this returns 11111111111111
	cnpj.Unmask("11.111.111/1111-11")

	// this returns 11.111.111/1111-11
	cnpj.Mask("11111111111111")
}
```

Based on [Go CPF](https://github.com/cuducos/go-cpf) ❤️
