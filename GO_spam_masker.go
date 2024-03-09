package main

import (
	"bufio"
	"fmt"
	"os"
	"sheremet-o/GO_spam_masker.git/masker"
)

func main() {
	message := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	words, _ := message.ReadString('\n')
	maskedMessage := masker.Masker(words)
	fmt.Println(maskedMessage)
}
