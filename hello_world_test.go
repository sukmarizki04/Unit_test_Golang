package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAdi(t *testing.T) {
	uji := HelloWorld(" Sukma")

	if uji != "Hello Sukma" {
		//error dan melanjutkan kebawah
		t.Error("Harusnya : He Sukma")
	}
	fmt.Println("Dieksekusi")
}

func TestHelloWorldSukma(t *testing.T) {
	uji := HelloWorld(" Sukma")

	if uji != "Hello Sukma" {
		//sedangkan testfailnow akan berhenti ketika menemukan function FailNow
		t.FailNow()
	}

	fmt.Println("TestHelloWorldSukma Done")
}
