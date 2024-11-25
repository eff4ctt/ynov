package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
    "os"
)

// Function to generate a random initialization vector (IV)
func generateIV() ([]byte, error) {
    iv := make([]byte, aes.BlockSize)
    if _, err := rand.Read(iv); err != nil {
        return nil, err
    }
    return iv, nil
}

// Function to encrypt the data
func encryptFile(filePath, key string) error {
    // Open the file for reading
    inputFile, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open input file: %v", err)
    }
    defer inputFile.Close()

    // Create an output file
    outputFile, err := os.Create(filePath + ".enc")
    if err != nil {
        return fmt.Errorf("failed to create output file: %v", err)
    }
    defer outputFile.Close()

    // Generate the AES cipher block
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return fmt.Errorf("failed to create AES cipher: %v", err)
    }

    // Generate an initialization vector (IV)
    iv, err := generateIV()
    if err != nil {
        return fmt.Errorf("failed to generate IV: %v", err)
    }

    // Write the IV to the output file (first 16 bytes are the IV)
    if _, err := outputFile.Write(iv); err != nil {
        return fmt.Errorf("failed to write IV to output file: %v", err)
    }

    // Create a cipher stream
    stream := cipher.NewCFBEncrypter(block, iv)

    // Read the file and encrypt it, writing the result to the output file
    buffer := make([]byte, 1024)
    for {
        bytesRead, err := inputFile.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            return fmt.Errorf("failed to read input file: %v", err)
        }

        // Encrypt the data
        stream.XORKeyStream(buffer[:bytesRead], buffer[:bytesRead])

        // Write the encrypted data to the output file
        if _, err := outputFile.Write(buffer[:bytesRead]); err != nil {
            return fmt.Errorf("failed to write encrypted data to output file: %v", err)
        }
    }

    return nil
}

// Function to decrypt the data
func decryptFile(filePath, key string) error {
    // Open the encrypted file for reading
    inputFile, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open input file: %v", err)
    }
    defer inputFile.Close()

    // Create an output file for the decrypted content
    outputFile, err := os.Create(filePath + ".dec")
    if err != nil {
        return fmt.Errorf("failed to create output file: %v", err)
    }
    defer outputFile.Close()

    // Generate the AES cipher block
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return fmt.Errorf("failed to create AES cipher: %v", err)
    }

    // Read the IV from the encrypted file
    iv := make([]byte, aes.BlockSize)
    if _, err := inputFile.Read(iv); err != nil {
        return fmt.Errorf("failed to read IV from file: %v", err)
    }

    // Create a cipher stream for decryption
    stream := cipher.NewCFBDecrypter(block, iv)

    // Read the encrypted file and decrypt it, writing the result to the output file
    buffer := make([]byte, 1024)
    for {
        bytesRead, err := inputFile.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            return fmt.Errorf("failed to read encrypted file: %v", err)
        }

        // Decrypt the data
        stream.XORKeyStream(buffer[:bytesRead], buffer[:bytesRead])

        // Write the decrypted data to the output file
        if _, err := outputFile.Write(buffer[:bytesRead]); err != nil {
            return fmt.Errorf("failed to write decrypted data to output file: %v", err)
        }
    }

    return nil
}

func main() {
    // Input file path and key for encryption/decryption
    filePath := "sample.txt" // Path to the file you want to encrypt
    key := "your-32-byte-long-encryption-key-here!" // AES key (must be 16, 24, or 32 bytes)

    // Encrypt the file
    err := encryptFile(filePath, key)
    if err != nil {
        fmt.Println("Error encrypting file:", err)
        return
    }
    fmt.Println("File encrypted successfully!")

    // Decrypt the file (for testing)
    err = decryptFile(filePath+".enc", key)
    if err != nil {
        fmt.Println("Error decrypting file:", err)
        return
    }
    fmt.Println("File decrypted successfully!")
}
