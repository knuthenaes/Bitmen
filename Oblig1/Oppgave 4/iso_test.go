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

// Under er en test av testen. Koden henter ascii2 (x80\x81\x82\x83\x84\x85) som beviser at testen fungerer
// dersom all kode er extended ascii code.

//func TestExtendedASCIIText(t *testing.T) {
//	for i := 0; i<len(ascii2);i++ {
//		if (ascii2[i] < 128) {
//			t.Fail()
//			fmt.Println("ERROR, value not in extended ASCII")
//		}
//	}
//}
