package encryption

import (
	"fmt"
	"os"
	"strings"
)

func GenerateKey(outputDir string) error {

	var encryptionKeyRaw []string

	for key, value := range CharMap {
		cryptoElement := fmt.Sprintf("%d=%s", value, key)
		encryptionKeyRaw = append(encryptionKeyRaw, cryptoElement)
	}

	finalEncryptionKey := strings.Join(encryptionKeyRaw, ",")

	keyfile, err := os.Create(fmt.Sprintf("%s/.key", outputDir))
	if err != nil {
		return err
	}
	_, errWrite := keyfile.WriteString(finalEncryptionKey)
	if errWrite != nil {
		return errWrite
	}

	return nil
}