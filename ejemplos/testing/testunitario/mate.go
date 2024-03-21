package testunitario

import "testing"

// func TestSuma(t *testing.T) {
// 	total := Suma(5, 5)

// 	if total != 10 {
// 		t.Errorf("Suma incorrecta %d se esperaba %d", total, 10)
// 	}
// }


func TestSuma(t *testing.T) {
	tabla := [] struct {
		a int
		b int
		c int
	}{
		{1,2,3},
		{2,8,4},
		{47,8,5},
	}

	total := Suma(5, 5)

	if total != 10 {
		t.Errorf("Suma incorrecta %d se esperaba %d", total, 10)
	}
}