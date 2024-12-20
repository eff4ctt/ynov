package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Decrypts file content
func decryptFile(filePath string, key []byte) error {
	// Read encrypted file content
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil // Ignore error, continue silently
	}

	// Check if the file has at least 12 bytes for the nonce
	if len(data) < 12 {
		return nil // Skip the file as it doesn't have enough data for decryption
	}

	// Extract the nonce and ciphertext
	nonce := data[:12]      // The first 12 bytes are the nonce
	ciphertext := data[12:] // The rest is the ciphertext

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil // Ignore error, continue silently
	}

	// AES-GCM decryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil // Ignore error, continue silently
	}

	// Decrypt the data using the nonce and ciphertext
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil // Ignore error, continue silently
	}

	// Save decrypted content back to the file
	err = ioutil.WriteFile(filePath, plaintext, 0644)
	if err != nil {
		return nil // Ignore error, continue silently
	}

	return nil
}

// Walk through the folder and decrypt all files
func decryptFilesInFolder(folderPath string, key []byte) error {
	return filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Ignore error, continue silently
		}
		if !info.IsDir() {
			// Try to decrypt the file and silently continue if there's an error
			err := decryptFile(path, key)
			if err != nil {
				return nil // Continue processing other files silently
			}
			// Optionally, you could still print successful decryption
			// fmt.Printf("Decrypted: %s\n", path)
		}
		return nil
	})
}

// Convert a hex string back to bytes
func hexToBytes(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

func main() {
	// Define the folder path to decrypt (change this as needed)
	folderPath := "C:/Users/hugo/Downloads/test" // Replace with the folder path you want to decrypt

	// Ask user for the decryption key (in hexadecimal format)
	var hexKey string
	fmt.Print("Enter the decryption key (in hexadecimal format): ")
	fmt.Scanln(&hexKey)

	// Convert the hex string back to the AES key
	key, err := hexToBytes(hexKey)
	if err != nil {
		// If you don't want to print anything here, just silently return
		return
	}

	// Decrypt all files in the folder
	err = decryptFilesInFolder(folderPath, key)
	if err != nil {
		// If you don't want to print anything here, just silently return
		return
	}

	// Optionally, you could print a message when decryption is complete
	// fmt.Println("Decryption complete.")
}
