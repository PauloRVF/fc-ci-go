package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(1, 1)
	expected := 2

	if total != expected {
		t.Errorf("Resultado da soma é invalido: %d esperado: %d", total, expected)
	}
}
