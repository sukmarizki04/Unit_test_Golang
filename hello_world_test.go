package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	uji := HelloWorld(" Sukma")

	if uji != "Hello Sukma" {
		fmt.Println("Tes Sukses")
		panic("HASIL TIDAK SESUAI")
	}
}

func TestHelloWorldSukma(t *testing.T) {
	uji := HelloWorld("Sukma")

	if uji != "Hello Sukma" {
		panic("HASIL TIDAK SESUAI")
	}
}
