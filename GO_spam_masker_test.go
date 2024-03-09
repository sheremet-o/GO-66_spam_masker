package main

import (
	"testing"
)

func Masker(message string) string {
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

func TestMasker(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Here's my spammy page: http://hehefouls.netHAHAHA see you.", "Here's my spammy page: http://******************* see you."},
		{"Перейдите по ссылке https://www.google.com для получения информации", "Перейдите по ссылке https://www.google.com для получения информации"},
		{"http://test.com is a test website", "http://******** is a test website"},
	}

	for _, testCase := range testCases {
		result := Masker(testCase.input)
		if result != testCase.expected {
			t.Errorf("Должно быть %s, а получаем %s", testCase.expected, result)
		}
	}
}
