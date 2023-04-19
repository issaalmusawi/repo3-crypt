package mycrypt

import (
	"fmt"
)

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

func Krypter(inputMessage []rune, chiffer int) ([]rune, error) {
	result := make([]rune, len(inputMessage))
	for i, r := range inputMessage {
		idx := sokIAlfabetet(r, ALF_SEM03)
		if idx == -1 {
			return nil, fmt.Errorf("invalid character: '%c'", r)
		}
		nyttIndeks := (idx + chiffer) % len(ALF_SEM03)
		result[i] = ALF_SEM03[nyttIndeks]
	}
	return result, nil
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
		}
	}
	return -1
}
