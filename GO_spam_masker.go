package main

import (
	"fmt"
	"os"
	"sheremet-o/GO_spam_masker.git/masker"
	"sheremet-o/GO_spam_masker.git/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите путь к файлу в аргументах запуска")
		return
	}

	filePath := os.Args[1]

	producer := &service.FileProducer{FilePath: filePath}
	presenter := &service.FileWriterPresenter{FilePath: "output.txt"}

	service := masker.NewMaskingService(producer, presenter)
	service.Run()
}
