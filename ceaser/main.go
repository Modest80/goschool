// main.go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func crypt(message string, key int) string {
	if key == 0 {
		return message
	}
	cryptMessage := bytes.NewBufferString("")
	for _, ch := range message {
		cryptMessage.WriteRune(rune(int(ch) + key))
	}
	return cryptMessage.String()
}

func input() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func main() {
	message := input()
	key := 2

	resultCrypt := crypt(message, key)
	fmt.Println(resultCrypt)
	resultDecrypt := crypt(resultCrypt, -key)
	fmt.Println(resultDecrypt)
}
