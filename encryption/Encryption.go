package encryption

import (
	"fmt"
	"math/rand"
)

func EncryptData(data string) ([]string, error) {

	for index, _ := range CharMap {
		randomNum := rand.Intn(1000)
		CharMap[index] = randomNum
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
