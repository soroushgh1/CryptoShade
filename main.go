package main

import (
	"CryptoShade/decryption"
	"CryptoShade/encryption"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//test.MainTest() // Developer test function, nothing important.

	args := os.Args

	switch args[1] {
	case "-e":
		var encryptedData string

		fmt.Println("Enter the data you want to encrypt: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		encryptedData = scanner.Text()

		var keyDir string

		fmt.Print("Enter what directory you want for the encryption key to be in ( . for current directory ) : ")

		scanner.Scan()
		keyDir = scanner.Text()

		encryptedList, err := encryption.EncryptData(encryptedData)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		encryptedDataString := strings.Join(encryptedList, "/")

		fmt.Println("Encrypted Data: ", encryptedDataString)

		keygenErr := encryption.GenerateKey(keyDir)
		if keygenErr != nil {
			log.Fatalln(keygenErr.Error())
			return
		}
		fmt.Println("Key generated !")
	case "-d":
		fmt.Print("Enter the text you want to encrypt: ")

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter the text you want to decrypt: ")
		decryptionInput, _ := reader.ReadString('\n')
		decryptionInput = strings.TrimSpace(decryptionInput)

		fmt.Print("Enter the path for your key: ")
		decryptionKeyDir, _ := reader.ReadString('\n')
		decryptionKeyDir = strings.TrimSpace(decryptionKeyDir)

		setKeyErr := decryption.SetMap(decryptionKeyDir)
		if setKeyErr != nil {
			log.Fatalln(setKeyErr.Error())
			return
		}

		decryptedList, err := decryption.DecryptData(decryptionInput)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		finalDecrypt := strings.Join(decryptedList, "")

		fmt.Println("Decrypted Data is : ", finalDecrypt)
	case "-h":
		var helpMessage string = `

	CryptoShade: One time text encryption Program
	

	Usage:
		cshade -e 		
		Prompts the user for encryption.
		

		cshade -d
		Prompts the user for decryption.
	
	More Info:
		Github Repository: https://www.github.com/soroushgh1/CryptoShade
	
	Created by Soroush GH in Go programing language.
 	
	`
	fmt.Println(helpMessage)
	}
}
