package q5

import (
	"strings"
	"unicode"
)

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {

	// Converter todas as letras maiúsculas para minúsculas
	s = strings.ToLower(s)

	// Remover caracteres não alfanuméricos
	var sb strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			sb.WriteRune(char)
		}
	}
	str := sb.String()

	// Comparar os caracteres da esquerda para a direita e da direita para a esquerda
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
