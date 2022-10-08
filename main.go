package main

import (
	"fmt"
	"strings"
)

func caesarCipherShiftRune(r rune, shift uint) rune {
	// mengambil shift dengan modulus 26
	// memakai 26 sesuai dengan jumlah alpabet
	s := rune(shift % 26)

	// jika shift = 0 maka akan langsung return kata tersebut
	if s == 0 {
		return r
	}

	// melakukan kalkulasi dengan menambah shift
	return (((r - 'A') + s) % 26) + 'A'
}

func caesarCipherEncrypt(value string, shift uint) string {
	var builder strings.Builder
	// melakukan perubahan menjadi huruf besar
	value = strings.ToUpper(value)

	for _, r := range value {
		// melakukan pengkondisian agar hanya abjad saja yang dilakukan enkripsi
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherShiftRune(r, shift)
		}

		builder.WriteRune(r)
	}

	return builder.String()
}

func main() {
	const encryptShift = 2
	const input = "A"

	encrypted := caesarCipherEncrypt(input, encryptShift)

	fmt.Printf(
		"INPUT: %s\nENCRYPTED: %s\n",
		input, encrypted,
	)
}
