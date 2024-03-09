package main

import "testing"

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
		result := masker(testCase.input)
		if result != testCase.expected {
			t.Errorf("Должно быть %s, а получаем %s", testCase.expected, result)
		}
	}
}
