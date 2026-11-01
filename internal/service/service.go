package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var ErrEmptyData = errors.New("service: пустые входные данные")

const morseSymbols = ".-/ \t\r\n"

func Convert(data string) (string, error) {
	s := strings.TrimSpace(data)
	if s == "" {
		return "", ErrEmptyData
	}

	if isMorse(s) {
		return morse.ToText(s), nil
	}

	return morse.ToMorse(s), nil
}

func isMorse(s string) bool {
	return strings.Trim(s, morseSymbols) == ""
}
