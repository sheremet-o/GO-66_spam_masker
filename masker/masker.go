package masker

import (
	"sheremet-o/GO_spam_masker.git/service"
)

type MaskingService struct {
	producer  service.Producer
	presenter service.Presenter
}

func NewMaskingService(producer service.Producer, presenter service.Presenter) *MaskingService {
	return &MaskingService{
		producer:  producer,
		presenter: presenter,
	}
}

func (ms *MaskingService) Run() {
	data, err := ms.producer.Produce()
	if err != nil {
		panic(err)
	}

	maskedData := make([]string, 0)
	for _, message := range data {
		maskedMessage := ms.Masker(message)
		maskedData = append(maskedData, maskedMessage)
	}

	err = ms.presenter.Present(maskedData)
	if err != nil {
		panic(err)
	}
}

func (ms *MaskingService) Masker(message string) string {
	buffer := []byte(message)
	linkHttp := []byte("http://")

	for i := 0; i < len(buffer)-len(linkHttp); i++ {
		if string(buffer[i:i+len(linkHttp)]) == string(linkHttp) {
			j := i + len(linkHttp)
			for j < len(buffer) && buffer[j] != ' ' {
				buffer[j] = '*'
				j++
			}
			i = j
		}
	}
	return string(buffer)
}
