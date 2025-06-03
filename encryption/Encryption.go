package encryption

import (
	"errors"
	"fmt"
	"math/rand"
)

func GenerateUniqueValue(existing map[string]int64) (int64, error) {

	used := make(map[int64]bool)
	for _, v := range existing {
		used[v] = true
	}

	const maxAttempts = 1000
	for i := 0; i < maxAttempts; i++ {
		val := rand.Intn(1000) + 1
		if !used[int64(val)] {
			return int64(val), nil
		}
	}

	return 0, errors.New("could not generate a unique value between 1–1000")
}

func EncryptData(data string) ([]string, error) {

	for index, _ := range CharMap {
		randomNum, Randerr := GenerateUniqueValue(CharMap)
		if Randerr != nil {
			return nil, Randerr
		}
		CharMap[index] = int64(randomNum)
	}

	var encrypted []string

	for i, letter := range data {
		num := CharMap[string(letter)]

		encryptedLetter := fmt.Sprintf("%d", num)

		if i == 0 {
			encryptedLetter = fmt.Sprintf("%d", num)
		}
		if i == len(data)-1 {
			encryptedLetter = fmt.Sprintf("%d", num)
		}
		encrypted = append(encrypted, encryptedLetter)
	}

	return encrypted, nil
}