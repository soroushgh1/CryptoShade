package decryption

import (
	"CryptoShade/encryption"
	"strconv"
	"strings"
)

func DecryptData(data string) ([]string, error) {

	splitedData := strings.Split(data, "/")

	var decryptedData []string

	for _, encryptedValue := range splitedData {

		encryptedValueNum, err := strconv.ParseInt(encryptedValue, 0, 64)
		if err != nil {
			return []string{}, err
		}

		for k, v := range encryption.CharMap {
			if v == encryptedValueNum {
				decryptedData = append(decryptedData, k)
				break
			}
		}

	}
	return decryptedData, nil
}