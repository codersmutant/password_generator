package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomPassword(character int) string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special := "!@#$%^&*~+-_"
	numbers := "1234567890"

	combine := letters + uppercaseLetter + special + numbers
	var password []byte

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	password = append(password, letters[rand.Intn(len(letters))])
	password = append(password, uppercaseLetter[rand.Intn(len(uppercaseLetter))])
	password = append(password, special[rand.Intn(len(special))])
	password = append(password, numbers[rand.Intn(len(numbers))])

	for i := 4; i < character; i++ {
		password = append(password, combine[rand.Intn(len(combine))])
	}

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func main() {
	fmt.Println("Pass generated:")

	fmt.Println(randomPassword(12))
}
