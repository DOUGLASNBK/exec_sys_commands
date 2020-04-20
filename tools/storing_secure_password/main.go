package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

var secretKey = "gamminogamma" //deprecate--- must be placed to a secure place..
var hashedPassword = "fff4f0d581e3bbcf3f8c944ba24a6932a23a0619314c63fed9ab4d482ef81411"

func GenerateSalt() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randomBytes), nil
}
func HashPassword(plainText string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, plainText)
	hashedValue := hash.Sum(nil)
	return hex.EncodeToString(hashedValue)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "fff4f0d581e3bbcf3f8c944ba24a6932a23a0619314c63fed9ab4d482ef81411" { // exit only if hash-verified passoword
		fmt.Print("Enter your password: ")
		scanner.Scan()
		text = scanner.Text()
		text = HashPassword(text)
		if text != "fff4f0d581e3bbcf3f8c944ba24a6932a23a0619314c63fed9ab4d482ef81411" {
			fmt.Println("Your hashed password was: ", text)
		}
	}

}