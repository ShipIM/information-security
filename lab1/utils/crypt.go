package utils

import (
	"errors"
	"strings"
)

var alphabet = []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя0123456789.,!?;:-()@#$%^&*_+=/\\|[]{}<> \n\r\t")

// VigenereEncrypt шифрует текст с помощью метода шифра Виженера.
func VigenereEncrypt(plaintext string, key []string) (string, error) {
	var encrypted strings.Builder
	keysValues, err := calculateKeysValues(key) // Получаем числовые значения символов ключа
	if err != nil {
		return "", err
	}

	i := 0 // Индекс для отслеживания позиции
	for _, character := range plaintext {
		plainIndex := indexOfRune(alphabet, character) // Находим индекс символа в открытом тексте
		if plainIndex == -1 {
			return "", errors.New("буква не найдена в алфавите: " + string(character))
		}

		keyIndex := plainIndex
		// Проходим по значениям ключа для вычисления нового индекса символа
		for _, values := range keysValues {
			keyIndex += values[i%(len(values))] // Используем модуль для циклического применения ключа
		}

		encryptedIndex := keyIndex % len(alphabet)
		encrypted.WriteRune(alphabet[encryptedIndex]) // Добавляем зашифрованный символ

		i++
	}

	return encrypted.String(), nil
}

// VigenereDecrypt расшифровывает зашифрованный текст с помощью метода шифра Виженера.
func VigenereDecrypt(encrypted string, key []string) (string, error) {
	var decrypted strings.Builder
	keysValues, err := calculateKeysValues(key) // Получаем числовые значения символов ключа
	if err != nil {
		return "", err
	}

	i := 0 // Индекс для отслеживания позиции в ключе
	for _, character := range encrypted {
		encryptedIndex := indexOfRune(alphabet, character) // Находим индекс зашифрованного символа
		if encryptedIndex == -1 {
			return "", errors.New("буква не найдена в алфавите: " + string(character))
		}

		keyIndex := encryptedIndex
		// Проходим по значениям ключа для вычисления оригинального индекса символа
		for _, values := range keysValues {
			keyIndex -= values[i%(len(values))] // Используем модуль для циклического применения ключа
			if keyIndex < 0 {
				keyIndex += len(alphabet) // Оборачиваем, если индекс стал отрицательным
			}
		}

		decryptedIndex := keyIndex % len(alphabet)
		decrypted.WriteRune(alphabet[decryptedIndex]) // Добавляем расшифрованный символ

		i++
	}

	return decrypted.String(), nil
}

// calculateKeysValues преобразует строки ключа в соответствующие числовые значения.
func calculateKeysValues(keys []string) ([][]int, error) {
	var keyIndices [][]int
	for _, key := range keys {
		var keyValues []int
		for _, character := range key {
			index := indexOfRune(alphabet, character) // Находим индекс символа ключа
			if index == -1 {
				return nil, errors.New("буква не найдена в алфавите: " + string(character))
			}
			keyValues = append(keyValues, index) // Сохраняем индекс
		}
		keyIndices = append(keyIndices, keyValues) // Добавляем значения для текущего ключа
	}

	return keyIndices, nil
}

// indexOfRune возвращает индекс руны в срезе рун, или -1, если не найдено.
func indexOfRune(runes []rune, r rune) int {
	for i, v := range runes {
		if v == r {
			return i
		}
	}
	return -1
}
