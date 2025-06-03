package decryption

import (
	"CryptoShade/encryption"
	"os"
	"strconv"
	"strings"
)

func SetMap(keydir string) error {

	encryptionKey, readerr := os.ReadFile(keydir)
	if readerr != nil {
		return readerr
	}

	splitedKey := strings.Split(string(encryptionKey), ",,")

	for _, keyValue := range splitedKey {
		splitedKeyValue := strings.Split(keyValue, "=")
		key := splitedKeyValue[0]
		value := splitedKeyValue[1]
		parseValue, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return err
		}
		encryption.CharMap[key] = parseValue
	}
	return nil
}