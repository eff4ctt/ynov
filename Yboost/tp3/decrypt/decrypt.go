package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Decrypt the given data using AES-GCM
func decryptData(data []byte, key []byte) ([]byte, error) {
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new GCM cipher for AES
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// The nonce size for GCM encryption
	nonceSize := gcm.NonceSize()

	// Ensure the data is long enough to contain a nonce + encrypted data
	if len(data) < nonceSize {
		return nil, fmt.Errorf("data too short to contain nonce")
	}

	// Extract the nonce from the data (first `nonceSize` bytes)
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt the ciphertext
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Decrypt files in the given folder and its subfolders
func decryptFilesInFolder(folderPath string, key []byte) error {
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error reading file:", err)
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Read file content
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("Error reading file:", err)
			return err
		}

		// Decrypt the file content
		plaintext, err := decryptData(data, key)
		if err != nil {
			log.Println("Error decrypting file:", err)
			return err
		}

		// Save the decrypted content back to the file
		err = ioutil.WriteFile(path, plaintext, info.Mode())
		if err != nil {
			log.Println("Error writing decrypted file:", err)
			return err
		}

		log.Println("Decrypted file:", path)
		return nil
	})

	return err
}

// Prompt user for the decryption key and validate the input
func promptForKey() ([]byte, error) {
	var keyHex string
	fmt.Print("Enter the encryption key (hex format): ")
	_, err := fmt.Scanln(&keyHex)
	if err != nil {
		return nil, err
	}

	// Decode the hex-encoded key
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid key format")
	}

	// Ensure the key length matches the required AES-256 size (32 bytes)
	if len(key) != 32 {
		return nil, fmt.Errorf("key must be 32 bytes for AES-256")
	}

	return key, nil
}

func main() {
	// Path to the folder to decrypt (change this to the folder you want to decrypt)
	folderPath := `C:\Users\hugo\Desktop\test\test.txt` // Update to your encrypted folder path

	// Prompt for the decryption key
	key, err := promptForKey()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Decrypt all files in the folder
	err = decryptFilesInFolder(folderPath, key)
	if err != nil {
		log.Fatalf("Error decrypting files: %v", err)
	}

	log.Println("Decryption process completed.")
}
