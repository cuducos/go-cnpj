package cnpj

import (
	"fmt"
	"testing"
)

func TestMask(t *testing.T) {
	for _, tc := range []struct {
		cnpj     string
		expected string
	}{
		{"11111111111111", "11.111.111/1111-11"},
		{"123456", "123456"},
		{"11223344556677889900", "11223344556677889900"},
		{"12ABC34501DE35", "12.ABC.345/01DE-35"},
	} {
		if got := Mask(tc.cnpj); tc.expected != got {
			t.Errorf("Mask(\"%s\") = %v; expected %s", tc.cnpj, got, tc.expected)
		}
	}
}

func TestUnmask(t *testing.T) {
	for _, tc := range []struct {
		cnpj     string
		expected string
	}{
		{"11.111.111/1111-11", "11111111111111"},
		{"12.ABC.345/01DE-35", "12ABC34501DE35"},
	} {
		if got := Unmask(tc.cnpj); tc.expected != got {
			t.Errorf("Unmask(\"%s\") = %v; want %s", tc.cnpj, got, tc.expected)
		}
	}
}

func TestIsValid(t *testing.T) {
	for _, tc := range []struct {
		cnpj     string
		expected bool
	}{
		{"11222333000181", true},
		{"11.222.333/0001-81", true},
		{"123", false},
		{"11.111.111/1111-11", false},
		{"12.ABC.345/01DE-35", true},
		{"12ABC34501DE35", true},
		{"12ABC34501DE42", false},
		{"12.345.678 9012-34", false},
		{"AB.CDE.FGH/IJKL-MN", false},
	} {
		if got := IsValid(tc.cnpj); tc.expected != got {
			t.Errorf("IsValid(\"%v\") = %v; expected %v", tc.cnpj, got, tc.expected)
		}
	}
}

func TestBase(t *testing.T) {
	for _, tc := range []struct {
		cnpj     string
		expected string
	}{
		{"11222333000181", "11222333"},
		{"11.222.333/0001-81", "11222333"},
		{"123", ""},
	} {
		if got := Base(tc.cnpj); tc.expected != got {
			t.Errorf("Base(%s) = %s; expected %s", tc.cnpj, got, tc.expected)
		}

	}
}

func TestOrder(t *testing.T) {
	for _, tc := range []struct {
		cnpj     string
		expected string
	}{
		{"11222333000181", "0001"},
		{"11.222.333/0001-81", "0001"},
		{"123", ""},
	} {
		if got := Order(tc.cnpj); tc.expected != got {
			t.Errorf("Order(%s) = %s; expected %s", tc.cnpj, got, tc.expected)
		}

	}
}

func TestCheckDigit(t *testing.T) {
	for _, tc := range []struct {
		cnpj     string
		expected string
	}{
		{"11222333000181", "81"},
		{"11.222.333/0001-81", "81"},
		{"123", ""},
	} {
		if got := CheckDigit(tc.cnpj); tc.expected != got {
			t.Errorf("CheckDigit(%s) = %s; expected %s", tc.cnpj, got, tc.expected)
		}

	}
}

func ExampleIsValid_validUnmasked() {
	fmt.Println(IsValid("11222333000181"))
	// Output: true
}

func ExampleIsValid_validMasked() {
	fmt.Println(IsValid("11.222.333/0001-81"))
	// Output: true
}

func ExampleIsValid_invalid() {
	fmt.Println(IsValid("11.111.111/1111-11"))
	// Output: false
}

func ExampleMask_valid() {
	fmt.Println(Mask("11111111111111"))
	// Output: 11.111.111/1111-11
}

func ExampleMask_invalid() {
	fmt.Println(Mask("42"))
	// Output: 42
}

func ExampleUnmask() {
	fmt.Println(Unmask("11.111.111/1111-11"))
	// Output: 11111111111111
}

func ExampleBase_validMasked() {
	fmt.Println(Base("11.222.333/0001-81"))
	// Output: 11222333
}

func ExampleBase_validUnmasked() {
	fmt.Println(Base("11222333000181"))
	// Output: 11222333
}

func ExampleBase_invalid() {
	fmt.Println(Base("123"))
	// Output:
}

func ExampleOrder_validMasked() {
	fmt.Println(Order("11.222.333/0001-81"))
	// Output: 0001
}

func ExampleOrder_validUnmasked() {
	fmt.Println(Order("11222333000181"))
	// Output: 0001
}

func ExampleOrder_invalid() {
	fmt.Println(Order("123"))
	// Output:
}

func ExampleCheckDigit_validMasked() {
	fmt.Println(CheckDigit("11.222.333/0001-81"))
	// Output: 81
}

func ExampleCheckDigit_validUnmasked() {
	fmt.Println(CheckDigit("11222333000181"))
	// Output: 81
}

func ExampleCheckDigit_invalid() {
	fmt.Println(CheckDigit("123"))
	// Output:
}
