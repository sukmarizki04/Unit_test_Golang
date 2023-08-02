package helper

import (
	"fmt"
	"testing"
)

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		panic("Salah harusnya %2.f")
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luast : %2.f", kubus.Luas())

	if kubus.Keliling() != kubus.Volume() {
		fmt.Printf("Thats True")
	}
}
