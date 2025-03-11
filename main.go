package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	useNumbers := flag.Bool("n", false, "Incluir números na senha")
	useLetters := flag.Bool("l", false, "Incluir letras na senha")
	useSpecial := flag.Bool("s", false, "Incluir caracteres especiais na senha")
	charCount := flag.Int("c", 12, "Quantidade de caracteres na senha")

	flag.Parse()

	if !*useNumbers && !*useLetters && !*useSpecial {
		*useNumbers = true
		*useLetters = true
		*useSpecial = true
	}

	password := generatePassword(*useNumbers, *useLetters, *useSpecial, *charCount)
	fmt.Println(password)
}

func generatePassword(useNumbers, useLetters, useSpecial bool, length int) string {
	rand.Seed(time.Now().UnixNano())

	numbers := "0123456789"
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	charset := ""
	if useNumbers {
		charset += numbers
	}
	if useLetters {
		charset += letters
	}
	if useSpecial {
		charset += special
	}

	if charset == "" {
		charset = numbers + letters + special
	}

	var password strings.Builder
	charsetLength := len(charset)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(charsetLength)
		password.WriteByte(charset[randomIndex])
	}

	return password.String()
}
