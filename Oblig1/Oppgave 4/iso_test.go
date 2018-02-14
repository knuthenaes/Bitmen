package main

import (
	"testing"
	"fmt"
)

func TestExtendedASCIIText(t *testing.T) {
	for i := 0; i<len(ascii);i++ {
		if (ascii[i] < 128) {
			t.Fail()
			fmt.Println("ERROR, value not in extended ASCII")
		}
	}
}

// Under er en test av testen. Koden henter ascii2 (x8a\x8b\x8c) som beviser at testen fungerer
// dersom all kode er extended ascii code.

//func TestExtendedASCIIText(t *testing.T) {
//	for i := 0; i<len(ascii2);i++ {
//		if (ascii2[i] < 128) {
//			t.Fail()
//			fmt.Println("ERROR, value not in extended ASCII")
//		}
//	}
//}
