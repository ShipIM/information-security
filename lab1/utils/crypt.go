package utils

import "strings"

func VigenereEncrypt(text, key string) string {
	var result strings.Builder
	keyLength := len(key)

	for i, char := range text {
		if char >= 'A' && char <= 'Z' {
			shift := key[i%keyLength] - 'A'
			result.WriteRune(((char-'A')+rune(shift))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			shift := key[i%keyLength] - 'a'
			result.WriteRune(((char-'a')+rune(shift))%26 + 'a')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func VigenereDecrypt(text, key string) string {
	var result strings.Builder
	keyLength := len(key)

	for i, char := range text {
		if char >= 'A' && char <= 'Z' {
			shift := key[i%keyLength] - 'A'
			result.WriteRune(((char-'A')-rune(shift)+26)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			shift := key[i%keyLength] - 'a'
			result.WriteRune(((char-'a')-rune(shift)+26)%26 + 'a')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func ComposeKey(keys []string) string {
	return strings.Join(keys, "")
}
