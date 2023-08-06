package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {

	result := HelloWorld(" Sukma")
	var err = "Program not found"
	require.Equal(t, "Hello Sukma", result, err)
	fmt.Println("Assert Require Done")
}

func TestHelloWorldAssertion(t *testing.T) {

	result := HelloWorld(" Sukma")
	var err = "Program not found"
	assert.Equal(t, "Hello Sukma", result, err)
	fmt.Println("TestHelloWOrldAssertion Done")

}
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

	if uji == "Hello Sukma" {
		//sedangkan testfailnow akan berhenti ketika menemukan function FailNow
		t.Fatal("Result must Be 'He Sukma'")
	}

	fmt.Println("TestHelloWorldSukma Done")
}
