package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func Encryption(data string) ([]string, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data = scanner.Text()

	for index, _ := range AlphabetNumberMap {
		randomNum := rand.Intn(1000)
		AlphabetNumberMap[index] = randomNum
	}

	var encrypted []string

	for i, letter := range data {
		num := AlphabetNumberMap[string(letter)]

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