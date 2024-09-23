package utils

import (
	"errors"
	"strings"
)

var alphabet = []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ ")

func VigenereEncrypt(plaintext string, key []string) (string, error) {
	var encrypted strings.Builder
	keysValues, err := calculateKeysValues(key)
	if err != nil {
		return "", err
	}

	i := 0
	for _, character := range plaintext {
		plainIndex := indexOfRune(alphabet, character)
		if plainIndex == -1 {
			return "", errors.New("буква не найдена в алфавите: " + string(character))
		}

		keyIndex := plainIndex
		for _, values := range keysValues {
			keyIndex += values[i%(len(values))]
		}

		encryptedIndex := keyIndex % len(alphabet)
		encrypted.WriteRune(alphabet[encryptedIndex])

		i++
	}

	return encrypted.String(), nil
}

func VigenereDecrypt(encrypted string, key []string) (string, error) {
	var decrypted strings.Builder
	keysValues, err := calculateKeysValues(key)
	if err != nil {
		return "", err
	}

	i := 0
	for _, character := range encrypted {
		encryptedIndex := indexOfRune(alphabet, character)
		if encryptedIndex == -1 {
			return "", errors.New("буква не найдена в алфавите: " + string(character))
		}

		keyIndex := encryptedIndex
		for _, values := range keysValues {
			keyIndex -= values[i%(len(values))]
			if keyIndex < 0 {
				keyIndex += len(alphabet)
			}
		}

		decryptedIndex := keyIndex % len(alphabet)
		decrypted.WriteRune(alphabet[decryptedIndex])

		i++
	}

	return decrypted.String(), nil
}

func calculateKeysValues(keys []string) ([][]int, error) {
	var keyIndices [][]int
	for _, key := range keys {
		var keyValues []int
		for _, character := range key {
			index := indexOfRune(alphabet, character)
			if index == -1 {
				return nil, errors.New("буква не найдена в алфавите: " + string(character))
			}
			keyValues = append(keyValues, index)
		}
		keyIndices = append(keyIndices, keyValues)
	}

	return keyIndices, nil
}

func indexOfRune(runes []rune, r rune) int {
	for i, v := range runes {
		if v == r {
			return i
		}
	}
	return -1
}
