package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet      = "123456789"
	specialCharSet = "#%"
)

func RandomPassword(n int) string {
	password := generatePassword(n)

	return password
}

func RandomPasswordFixPosition() string {

	index0 := string(upperCharSet[rand.Intn(len(upperCharSet))])
	index1 := string(specialCharSet[rand.Intn(len(specialCharSet))])
	index2 := string(lowerCharSet[rand.Intn(len(lowerCharSet))])
	index3 := string(lowerCharSet[rand.Intn(len(lowerCharSet))])
	index4 := string(lowerCharSet[rand.Intn(len(lowerCharSet))])
	index5 := string(numberSet[rand.Intn(len(numberSet))])
	index6 := string(numberSet[rand.Intn(len(numberSet))])
	index7 := string(numberSet[rand.Intn(len(numberSet))])
	index8 := string(numberSet[rand.Intn(len(numberSet))])
	index9 := string(upperCharSet[rand.Intn(len(upperCharSet))])

	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", index0, index1, index2, index3, index4, index5, index6, index7, index8, index9)
}

func GetUnix() string {
	timeUnix := time.Now().Unix()
	unix := strconv.Itoa(int(timeUnix))

	return unix
}

func generatePassword(passwordLength int) string {
	var password strings.Builder

	for i := 0; i < passwordLength; i++ {
		if i == 2 || i == 4 {
			random := rand.Intn(len(specialCharSet))
			password.WriteString(string(specialCharSet[random]))
		} else if i == 3 || i == 5 {
			random := rand.Intn(len(numberSet))
			password.WriteString(string(numberSet[random]))
		} else if i == 6 || i == 7 {
			random := rand.Intn(len(upperCharSet))
			password.WriteString(string(upperCharSet[random]))
		} else {
			random := rand.Intn(len(lowerCharSet))
			password.WriteString(string(lowerCharSet[random]))
		}
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
