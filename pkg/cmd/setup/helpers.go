package setup

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatePassword() (string, error) {
	s := make([]byte, 64)
	for i := range s {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		s[i] = charset[idx.Int64()]
	}
	return string(s), nil
}

func prompt(reader *bufio.Reader, defaultValue string, message string) (string, error) {
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if text == "\n" {
		return defaultValue, nil
	}

	return strings.TrimSpace(text), nil
}

func promptPassword(reader *bufio.Reader, message string) (string, error) {
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if text == "\n" {
		password, err := generatePassword()
		if err != nil {
			return "", err
		}

		return password, nil
	}

	return strings.TrimSpace(text), nil
}

func promptDuration(reader *bufio.Reader, defaultValue time.Duration, minValue time.Duration, message string) (time.Duration, error) {
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	if text == "\n" {
		return defaultValue, nil
	}

	out, err := time.ParseDuration(strings.TrimSpace(text))
	if err != nil {
		return 0, err
	}

	if out < minValue {
		return 0, fmt.Errorf("duration must be greater than %s", minValue)
	}

	return out, nil
}

func promptInt(reader *bufio.Reader, defaultValue int, message string) (int, error) {
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	if text == "\n" {
		return defaultValue, nil
	}

	return strconv.Atoi(strings.TrimSpace(text))
}
