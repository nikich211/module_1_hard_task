package generate_password

import (
	"fmt"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(n int) (string, error) {
	if n < 1 {
		return "", nil
	}
	runes := []rune(letters)
	var result []rune
	for i := 0; i < n; i++ {
		result = append(result, runes[rand.Intn(len(runes))])
	}
	return string(result), nil
}
func main() {
	str, err := GeneratePassword(8)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
